package benchserder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
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

func TestEvent_MarshalUnmarshal(t *testing.T) {
	for _, tc := range marshallers {
		t.Run(fmt.Sprintf("marshaller=%s", tc.name), func(t *testing.T) {
			for _, tf := range testEvents {
				t.Run(fmt.Sprintf("file=%s", tf), func(t *testing.T) {
					data := mustReadFile(t, tf)
					event := &Event{}
					require.NoError(t, json.Unmarshal(data, &event))

					gotData, err := tc.marshaller.Marshal(event)
					require.NoError(t, err)
					require.NotEmpty(t, gotData)

					gotEvent := &Event{}
					require.NoError(t, tc.marshaller.Unmarshal(gotData, gotEvent))
					//assert.Equal(t, event, gotEvent)

					d1, _ := json.Marshal(event)
					d2, _ := json.Marshal(gotEvent)
					assert.JSONEq(t, string(d1), string(d2))
				})
			}
		})
	}
}

func BenchmarkEvent_Marshal(b *testing.B) {
	for _, tc := range marshallers {
		b.Run(fmt.Sprintf("marshaller=%s", tc.name), func(b *testing.B) {
			for _, tf := range testEvents {
				b.Run(fmt.Sprintf("file=%s", tf), func(b *testing.B) {
					data := mustReadFile(b, tf)
					event := &Event{}
					err := json.Unmarshal(data, event)
					if err != nil {
						b.Fatal(err)
					}
					benchmark_Marshal(b, tc.marshaller, event)
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
		return tm.Marshal(event)
	}

	for _, tc := range marshallers {
		b.Run(fmt.Sprintf("marshaller=%s", tc.name), func(b *testing.B) {
			for _, tf := range testEvents {
				b.Run(fmt.Sprintf("file=%s", tf), func(b *testing.B) {
					data, err := prepareData(tf, tc.marshaller)
					if err != nil {
						b.Fatal(err)
					}
					event := &Event{}
					benchmark_Unmarshal(b, tc.marshaller, data, event)
				})
			}
		})
	}
}

func BenchmarkEvent_Compression(b *testing.B) {
	for _, tc := range marshallers {
		b.Run(fmt.Sprintf("marshaller=%s", tc.name), func(b *testing.B) {
			for _, tf := range testEvents {
				b.Run(fmt.Sprintf("file=%s", tf), func(b *testing.B) {
					data := mustReadFile(b, tf)
					event := &Event{}
					err := json.Unmarshal(data, event)
					if err != nil {
						b.Fatal(err)
					}
					mData, err := tc.marshaller.Marshal(event)
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
		data, err := tm.Marshal(v)
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
		err := tm.Unmarshal(data, v)
		if err != nil {
			b.Fatal(err)
		}
		if v == nil {
			b.Fatal("unexpected nil value")
		}
	}
}

func readFile(tf string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join("testdata", tf))
}

func mustReadFile(t testing.TB, tf string) []byte {
	data, err := readFile(tf)
	if err != nil {
		t.Fatal(err)
	}
	return data
}
