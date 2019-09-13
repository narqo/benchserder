package benchserder

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testEvents = []string{
	"event_2K.json",
	"event_5K.json",
	"event_10K.json",
	"event_50K.json",
	"event_100K.json",
	"event_225K.json",
}

//func TestEvent_MarshalUnmarshalJSON(t *testing.T) {
//	for _, tf := range testEvents {
//		t.Run(tf, func(t *testing.T) {
//			data := mustReadFile(t, tf)
//
//			var event Event
//			require.NoError(t, json.Unmarshal(data, &event))
//
//			gotData, err := json.Marshal(event)
//			require.NoError(t, err)
//			assert.JSONEq(t, string(data), string(gotData))
//		})
//	}
//}
//
//func TestEvent_MarshalUnmarshalMsgpack(t *testing.T) {
//	for _, tf := range testEvents {
//		t.Run(tf, func(t *testing.T) {
//			data := mustReadFile(t, tf)
//			var event Event
//			require.NoError(t, json.Unmarshal(data, &event))
//
//			t.Run("as struct", func(t *testing.T) {
//				mData, err := marshalMsgpack(event)
//				require.NoError(t, err)
//
//				err = unmarshalMsgpack(mData, &event)
//				require.NoError(t, err)
//
//				gotData, _ := json.Marshal(event)
//				assert.JSONEq(t, string(data), string(gotData))
//			})
//
//			t.Run("as array", func(t *testing.T) {
//				mData, err := marshalMsgpackAsArray(event)
//				require.NoError(t, err)
//
//				err = unmarshalMsgpack(mData, &event)
//				require.NoError(t, err)
//
//				gotData, _ := json.Marshal(event)
//				assert.JSONEq(t, string(data), string(gotData))
//			})
//		})
//	}
//}
//
//func TestEvent_MarshalUnmarshalCodecCBOR(t *testing.T) {
//	for _, tf := range testEvents {
//		t.Run(tf, func(t *testing.T) {
//			data := mustReadFile(t, tf)
//			var event Event
//			require.NoError(t, json.Unmarshal(data, &event))
//
//			cData, err := marshalCodecCBOR(event)
//			require.NoError(t, err)
//
//			err = unmarshalCodecCBOR(cData, &event)
//			require.NoError(t, err)
//
//			gotData, _ := json.Marshal(event)
//			assert.JSONEq(t, string(data), string(gotData))
//		})
//	}
//}
//
//func TestEvent_MarshalUnmarshalFxamackerCBOR(t *testing.T) {
//	for _, tf := range testEvents {
//		t.Run(tf, func(t *testing.T) {
//			data := mustReadFile(t, tf)
//			var event Event
//			require.NoError(t, json.Unmarshal(data, &event))
//
//			cData, err := marshalFxamackerCBOR(event)
//			require.NoError(t, err)
//
//			err = unmarshalFxamackerCBOR(cData, &event)
//			require.NoError(t, err)
//
//			gotData, _ := json.Marshal(event)
//			assert.JSONEq(t, string(data), string(gotData))
//		})
//	}
//}
//
//func TestEvent_MarshalUnmarshalGogoProto(t *testing.T) {
//	for _, tf := range testEvents {
//		t.Run(tf, func(t *testing.T) {
//			data := mustReadFile(t, tf)
//			event := &Event{}
//			require.NoError(t, json.Unmarshal(data, event))
//
//			pData, err := marshalGogoProto(event)
//			require.NoError(t, err)
//
//			err = unmarshalGogoProto(pData, event)
//			require.NoError(t, err)
//
//			gotData, _ := json.Marshal(event)
//			assert.JSONEq(t, string(data), string(gotData))
//		})
//	}
//}

func BenchmarkEvent_Marshal(b *testing.B) {
	for _, tc := range marshallers {
		b.Run(fmt.Sprintf("marshaller=%s", tc.name), func(b *testing.B) {
			tm := tc.marshaller()
			for _, tf := range testEvents {
				b.Run(fmt.Sprintf("file=%s", tf), func(b *testing.B) {
					data := mustReadFile(b, tf)
					event := &Event{}
					err := json.Unmarshal(data, event)
					if err != nil {
						b.Fatal(err)
					}
					benchmark_Marshal(b, tm, event)
				})
			}
		})
	}
}

func BenchmarkEvent_Unmarshal(b *testing.B) {
	prepareData := func(tf string, tm *testMarshaller) ([]byte, error) {
		data := mustReadFile(b, tf)
		event := &Event{}
		if err := json.Unmarshal(data, event); err != nil {
			b.Fatal(err)
		}
		return tm.marshalFunc(event)
	}

	for _, tc := range marshallers {
		b.Run(fmt.Sprintf("marshaller=%s", tc.name), func(b *testing.B) {
			tm := tc.marshaller()
			for _, tf := range testEvents {
				b.Run(fmt.Sprintf("file=%s", tf), func(b *testing.B) {
					data, err := prepareData(tf, tm)
					if err != nil {
						b.Fatal(err)
					}
					event := &Event{}
					benchmark_Unmarshal(b, tm, data, event)
				})
			}
		})
	}
}

func BenchmarkEvent_Compression(b *testing.B) {
	for _, tc := range marshallers {
		b.Run(fmt.Sprintf("marshaller=%s", tc.name), func(b *testing.B) {
			tm := tc.marshaller()
			for _, tf := range testEvents {
				b.Run(fmt.Sprintf("file=%s", tf), func(b *testing.B) {
					data := mustReadFile(b, tf)
					event := &Event{}
					err := json.Unmarshal(data, event)
					if err != nil {
						b.Fatal(err)
					}
					mData, err := tm.marshalFunc(event)
					if err != nil {
						b.Fatal(err)
					}
					b.ReportMetric(0, "ns/op")
					b.ReportMetric(float64(len(mData)), "bytes")
					b.ReportMetric(float64(len(data))/float64(len(mData)), "%compression")
				})
			}
		})
	}
}

func benchmark_Marshal(b *testing.B, tm *testMarshaller, v interface{}) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := tm.marshalFunc(v)
		if err != nil {
			b.Fatal(err)
		}
		if data == nil {
			b.Fatal("unexpected nil data")
		}
	}
}

func benchmark_Unmarshal(b *testing.B, tm *testMarshaller, data []byte, v interface{}) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := tm.unmarshalFunc(data, v)
		if err != nil {
			b.Fatal(err)
		}
		if v == nil {
			b.Fatal("unexpected nil value")
		}
	}
}
