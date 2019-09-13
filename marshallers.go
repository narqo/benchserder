package benchserder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/fxamacker/cbor"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/ugorji/go/codec"
	"github.com/vmihailenco/msgpack/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type testMarshaller struct {
	marshalFunc   func(interface{}) ([]byte, error)
	unmarshalFunc func([]byte, interface{}) error
}

var marshallers = []struct {
	name       string
	marshaller func() *testMarshaller
}{
	{
		"json",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc:   json.Marshal,
				unmarshalFunc: json.Unmarshal,
			}
		},
	},
	{
		"easyjson",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc: func(v interface{}) ([]byte, error) {
					return easyjson.Marshal(v.(easyjson.Marshaler))
				},
				unmarshalFunc: func(data []byte, v interface{}) error {
					return easyjson.Unmarshal(data, v.(easyjson.Unmarshaler))
				},
			}
		},
	},
	{
		"jsoniter",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc:   jsoniter.ConfigCompatibleWithStandardLibrary.Marshal,
				unmarshalFunc: jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal,
			}
		},
	},
	{
		"vmihailenco-msgpack",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc: func(v interface{}) ([]byte, error) {
					var buf bytes.Buffer
					enc := msgpack.NewEncoder(&buf)
					enc.UseJSONTag(true) // must use json tags
					err := enc.Encode(v)
					return buf.Bytes(), err
				},
				unmarshalFunc: func(data []byte, v interface{}) error {
					dec := msgpack.NewDecoder(bytes.NewReader(data))
					dec.UseJSONTag(true)
					return dec.Decode(v)
				},
			}
		},
	},
	{
		"vmihailenco-msgpack-as-array",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc: func(v interface{}) ([]byte, error) {
					var buf bytes.Buffer
					enc := msgpack.NewEncoder(&buf)
					enc.StructAsArray(true) // as array
					err := enc.Encode(v)
					return buf.Bytes(), err
				},
				unmarshalFunc: func(data []byte, v interface{}) error {
					dec := msgpack.NewDecoder(bytes.NewReader(data))
					dec.UseJSONTag(true) // do we need json tags here?
					return dec.Decode(v)
				},
			}
		},
	},
	{
		"codec-json",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc: func(v interface{}) ([]byte, error) {
					var h codec.JsonHandle
					var buf bytes.Buffer
					err := codec.NewEncoder(&buf, &h).Encode(v)
					return buf.Bytes(), err
				},
				unmarshalFunc: func(data []byte, v interface{}) error {
					var h codec.JsonHandle
					return codec.NewDecoderBytes(data, &h).Decode(v)
				},
			}
		},
	},
	{
		"codec-msgpack",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc: func(v interface{}) ([]byte, error) {
					var h codec.MsgpackHandle
					var buf bytes.Buffer
					err := codec.NewEncoder(&buf, &h).Encode(v)
					return buf.Bytes(), err
				},
				unmarshalFunc: func(data []byte, v interface{}) error {
					var h codec.MsgpackHandle
					return codec.NewDecoderBytes(data, &h).Decode(v)
				},
			}
		},
	},
	{
		"codec-cbor",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc: func(v interface{}) ([]byte, error) {
					var h codec.CborHandle
					var buf bytes.Buffer
					err := codec.NewEncoder(&buf, &h).Encode(v)
					return buf.Bytes(), err
				},
				unmarshalFunc: func(data []byte, v interface{}) error {
					var h codec.CborHandle
					return codec.NewDecoderBytes(data, &h).Decode(v)
				},
			}
		},
	},
	{
		"fxamacker-cbor",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc: func(v interface{}) ([]byte, error) {
					return cbor.Marshal(v, cbor.EncOptions{Canonical: false})
				},
				unmarshalFunc: func(data []byte, v interface{}) error {
					return cbor.Unmarshal(data, v)
				},
			}
		},
	},
	{
		"bson",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc:   bson.Marshal,
				unmarshalFunc: bson.Unmarshal,
			}
		},
	},
	{
		"gogo-proto",
		func() *testMarshaller {
			return &testMarshaller{
				marshalFunc: func(v interface{}) ([]byte, error) {
					return v.(proto.Marshaler).Marshal()
				},
				unmarshalFunc: func(data []byte, v interface{}) error {
					return v.(proto.Unmarshaler).Unmarshal(data)
				},
			}
		},
	},
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
