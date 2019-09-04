package benchserder

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

var testEvents = []string{
	"event_2K.json",
	"event_10K.json",
	"event_25K.json",
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

			mData, err := marshalMsgpack(event)
			require.NoError(t, err)

			err = unmarshalMsgpack(mData, &event)
			require.NoError(t, err)

			gotData, _ := json.Marshal(event)
			assert.JSONEq(t, string(data), string(gotData))
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

func BenchmarkEvent_Marshal(b *testing.B) {
	cases := map[string]marshanFunc{
		"json":                json.Marshal,
		"easyjson":            marshalEasyJSON,
		"vmihailenco/msgpack": marshalMsgpack,
		"codec/msgpack":       marshalCodecMsgpack,
		"codec/cbor":          marshalCodecCBOR,
		"fxamacker/cbor":      marshalFxamackerCBOR,
		"bson":                marshalBSON,
	}

	for format, fn := range cases {
		b.Run(format, func(b *testing.B) {
			for _, tf := range testEvents {
				b.Run(tf, func(b *testing.B) {
					data := mustReadFile(b, tf)
					var event Event
					err := json.Unmarshal(data, &event)
					if err != nil {
						b.Fatal(err)
					}

					benchmark_Marshal(b, event, fn)
				})
			}
		})
	}
}

func BenchmarkEvent_Unmarshal(b *testing.B) {
	b.Run("json", func(b *testing.B) {
		for _, tf := range testEvents {
			b.Run(tf, func(b *testing.B) {
				data := mustReadFile(b, tf)
				var event Event
				benchmark_Unmarshal(b, data, &event, json.Unmarshal)
			})
		}
	})

	prepareData := func(tf string, fn marshanFunc) ([]byte, error) {
		data := mustReadFile(b, tf)
		var event Event
		if err := json.Unmarshal(data, &event); err != nil {
			b.Fatal(err)
		}

		return fn(event)
	}

	b.Run("easyjson", func(b *testing.B) {
		for _, tf := range testEvents {
			b.Run(tf, func(b *testing.B) {
				data := mustReadFile(b, tf)
				var event Event
				benchmark_Unmarshal(b, data, &event, unmarshalEasyJSON)
			})
		}
	})

	b.Run("vmihailenco/msgpack", func(b *testing.B) {
		for _, tf := range testEvents {
			b.Run(tf, func(b *testing.B) {
				data, err := prepareData(tf, marshalMsgpack)
				if err != nil {
					b.Fatal(err)
				}

				var event Event
				benchmark_Unmarshal(b, data, &event, unmarshalMsgpack)
			})
		}
	})

	b.Run("codec/msgpack", func(b *testing.B) {
		for _, tf := range testEvents {
			b.Run(tf, func(b *testing.B) {
				data, err := prepareData(tf, marshalCodecMsgpack)
				if err != nil {
					b.Fatal(err)
				}

				var event Event
				benchmark_Unmarshal(b, data, &event, unmarshalCodecMsgpack)
			})
		}
	})

	b.Run("codec/cbor", func(b *testing.B) {
		for _, tf := range testEvents {
			b.Run(tf, func(b *testing.B) {
				data, err := prepareData(tf, marshalCodecCBOR)
				if err != nil {
					b.Fatal(err)
				}

				var event Event
				benchmark_Unmarshal(b, data, &event, unmarshalCodecCBOR)
			})
		}
	})

	b.Run("fxamacker/cbor", func(b *testing.B) {
		for _, tf := range testEvents {
			b.Run(tf, func(b *testing.B) {
				data, err := prepareData(tf, marshalFxamackerCBOR)
				if err != nil {
					b.Fatal(err)
				}

				var event Event
				benchmark_Unmarshal(b, data, &event, unmarshalFxamackerCBOR)
			})
		}
	})

	b.Run("bson", func(b *testing.B) {
		for _, tf := range testEvents {
			b.Run(tf, func(b *testing.B) {
				data, err := prepareData(tf, bson.Marshal)
				if err != nil {
					b.Fatal(err)
				}

				var event Event
				benchmark_Unmarshal(b, data, &event, unmarshalBSON)
			})
		}
	})
}

func benchmark_Marshal(b *testing.B, v interface{}, fn marshanFunc) {
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
