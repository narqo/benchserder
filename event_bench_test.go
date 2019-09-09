package benchserder

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testEvents = []string{
	"event_2K.json",
	"event_5K.json",
	"event_10K.json",
	"event_50K.json",
	"event_100K.json",
	"event_225K.json",
}

func TestEvent_MarshalUnmarshalJSON(t *testing.T) {
	for _, tf := range testEvents {
		t.Run(tf, func(t *testing.T) {
			data := mustReadFile(t, tf)

			var event Event
			require.NoError(t, json.Unmarshal(data, &event))

			gotData, err := json.Marshal(event)
			require.NoError(t, err)
			assert.JSONEq(t, string(data), string(gotData))
		})
	}
}

func TestEvent_MarshalUnmarshalMsgpack(t *testing.T) {
	for _, tf := range testEvents {
		t.Run(tf, func(t *testing.T) {
			data := mustReadFile(t, tf)
			var event Event
			require.NoError(t, json.Unmarshal(data, &event))

			t.Run("as struct", func(t *testing.T) {
				mData, err := marshalMsgpack(event)
				require.NoError(t, err)

				err = unmarshalMsgpack(mData, &event)
				require.NoError(t, err)

				gotData, _ := json.Marshal(event)
				assert.JSONEq(t, string(data), string(gotData))
			})

			t.Run("as array", func(t *testing.T) {
				mData, err := marshalMsgpackAsArray(event)
				require.NoError(t, err)

				err = unmarshalMsgpack(mData, &event)
				require.NoError(t, err)

				gotData, _ := json.Marshal(event)
				assert.JSONEq(t, string(data), string(gotData))
			})
		})
	}
}

func TestEvent_MarshalUnmarshalCodecCBOR(t *testing.T) {
	for _, tf := range testEvents {
		t.Run(tf, func(t *testing.T) {
			data := mustReadFile(t, tf)
			var event Event
			require.NoError(t, json.Unmarshal(data, &event))

			cData, err := marshalCodecCBOR(event)
			require.NoError(t, err)

			err = unmarshalCodecCBOR(cData, &event)
			require.NoError(t, err)

			gotData, _ := json.Marshal(event)
			assert.JSONEq(t, string(data), string(gotData))
		})
	}
}

func TestEvent_MarshalUnmarshalFxamackerCBOR(t *testing.T) {
	for _, tf := range testEvents {
		t.Run(tf, func(t *testing.T) {
			data := mustReadFile(t, tf)
			var event Event
			require.NoError(t, json.Unmarshal(data, &event))

			cData, err := marshalFxamackerCBOR(event)
			require.NoError(t, err)

			err = unmarshalFxamackerCBOR(cData, &event)
			require.NoError(t, err)

			gotData, _ := json.Marshal(event)
			assert.JSONEq(t, string(data), string(gotData))
		})
	}
}

func TestEvent_MarshalUnmarshalGogoProto(t *testing.T) {
	for _, tf := range testEvents {
		t.Run(tf, func(t *testing.T) {
			data := mustReadFile(t, tf)
			event := &Event{}
			require.NoError(t, json.Unmarshal(data, event))

			pData, err := marshalGogoProto(event)
			require.NoError(t, err)

			err = unmarshalGogoProto(pData, event)
			require.NoError(t, err)

			gotData, _ := json.Marshal(event)
			assert.JSONEq(t, string(data), string(gotData))
		})
	}
}

var testCases = map[string]struct {
	marshalFunc   marshalFunc
	unmarshalFunc unmarshalFunc
}{
	"json":                    {json.Marshal, json.Unmarshal},
	"mailru-easyjson":         {marshalEasyJSON, unmarshalEasyJSON},
	"vmihailenco-msgpack":     {marshalMsgpack, unmarshalMsgpack},
	"vmihailenco-msgpack-arr": {marshalMsgpackAsArray, unmarshalMsgpack},
	"codec-msgpack":           {marshalCodecMsgpack, unmarshalCodecMsgpack},
	"codec-cbor":              {marshalCodecCBOR, unmarshalCodecCBOR},
	"fxamacker-cbor":          {marshalFxamackerCBOR, unmarshalFxamackerCBOR},
	"bson":                    {marshalBSON, unmarshalBSON},
	"gogo-proto":              {marshalGogoProto, unmarshalGogoProto},
}

func BenchmarkEvent_Marshal(b *testing.B) {
	for format, tc := range testCases {
		b.Run(format, func(b *testing.B) {
			for _, tf := range testEvents {
				b.Run(tf, func(b *testing.B) {
					data := mustReadFile(b, tf)
					event := &Event{}
					err := json.Unmarshal(data, event)
					if err != nil {
						b.Fatal(err)
					}

					benchmark_Marshal(b, event, tc.marshalFunc)
				})
			}
		})
	}
}

func BenchmarkEvent_Unmarshal(b *testing.B) {
	prepareData := func(tf string, fn marshalFunc) ([]byte, error) {
		data := mustReadFile(b, tf)
		event := &Event{}
		if err := json.Unmarshal(data, event); err != nil {
			b.Fatal(err)
		}
		return fn(event)
	}

	for format, tc := range testCases {
		b.Run(format, func(b *testing.B) {
			for _, tf := range testEvents {
				b.Run(tf, func(b *testing.B) {
					data, err := prepareData(tf, tc.marshalFunc)
					if err != nil {
						b.Fatal(err)
					}
					event := &Event{}
					benchmark_Unmarshal(b, data, event, tc.unmarshalFunc)
				})
			}
		})
	}
}

func benchmark_Marshal(b *testing.B, v interface{}, fn marshalFunc) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := fn(v)
		if err != nil {
			b.Fatal(err)
		}
		if data == nil {
			b.Fatal("unexpected nil data")
		}
	}
}

func benchmark_Unmarshal(b *testing.B, data []byte, v interface{}, fn unmarshalFunc) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := fn(data, v)
		if err != nil {
			b.Fatal(err)
		}
		if v == nil {
			b.Fatal("unexpected nil value")
		}
	}
}
