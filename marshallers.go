package benchserder

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"github.com/fxamacker/cbor"
	cborv2 "github.com/fxamacker/cbor/v2"
	goccyjson "github.com/goccy/go-json"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/pquerna/ffjson/ffjson"
	segmentiojson "github.com/segmentio/encoding/json"
	"github.com/ugorji/go/codec"
	"github.com/vmihailenco/msgpack/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var cborv2Encoder cborv2.EncMode

func init() {
	encOpts := cborv2.CanonicalEncOptions()
	encOpts.Time = cborv2.TimeUnixDynamic
	cborv2Encoder, _ = encOpts.EncMode()
}

type testMarshaller struct {
	Marshal   func(interface{}) ([]byte, error)
	Unmarshal func([]byte, interface{}) error
}

var marshallers = []struct {
	name       string
	marshaller *testMarshaller
}{
	{
		"gob",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				var buf bytes.Buffer
				err := gob.NewEncoder(&buf).Encode(v)
				return buf.Bytes(), err
			},
			Unmarshal: func(data []byte, v interface{}) error {
				return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
			},
		},
	},
	{
		"json",
		&testMarshaller{
			Marshal:   json.Marshal,
			Unmarshal: json.Unmarshal,
		},
	},
	{
		"easyjson",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				return easyjson.Marshal(v.(easyjson.Marshaler))
			},
			Unmarshal: func(data []byte, v interface{}) error {
				return easyjson.Unmarshal(data, v.(easyjson.Unmarshaler))
			},
		},
	},
	{
		"jsoniter",
		&testMarshaller{
			Marshal:   jsoniter.ConfigCompatibleWithStandardLibrary.Marshal,
			Unmarshal: jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal,
		},
	},
	{
		"ffjson",
		&testMarshaller{
			Marshal:   ffjson.MarshalFast,
			Unmarshal: ffjson.UnmarshalFast,
		},
	},
	{
		"segmentio-json",
		&testMarshaller{
			Marshal:   segmentiojson.Marshal,
			Unmarshal: segmentiojson.Unmarshal,
		},
	},
	{
		"goccy-json",
		&testMarshaller{
			Marshal:   goccyjson.Marshal,
			Unmarshal: goccyjson.Unmarshal,
		},
	},
	{
		"vmihailenco-msgpack",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				var buf bytes.Buffer
				enc := msgpack.NewEncoder(&buf)
				enc.UseJSONTag(true) // must use json tags
				err := enc.Encode(v)
				return buf.Bytes(), err
			},
			Unmarshal: func(data []byte, v interface{}) error {
				dec := msgpack.NewDecoder(bytes.NewReader(data))
				dec.UseJSONTag(true)
				return dec.Decode(v)
			},
		},
	},
	{
		"vmihailenco-msgpack-as-array",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				var buf bytes.Buffer
				enc := msgpack.NewEncoder(&buf)
				enc.StructAsArray(true) // as array
				err := enc.Encode(v)
				return buf.Bytes(), err
			},
			Unmarshal: func(data []byte, v interface{}) error {
				dec := msgpack.NewDecoder(bytes.NewReader(data))
				dec.UseJSONTag(true) // do we need json tags here?
				return dec.Decode(v)
			},
		},
	},
	{
		"codec-json",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				var h codec.JsonHandle
				var buf bytes.Buffer
				err := codec.NewEncoder(&buf, &h).Encode(v)
				return buf.Bytes(), err
			},
			Unmarshal: func(data []byte, v interface{}) error {
				var h codec.JsonHandle
				return codec.NewDecoderBytes(data, &h).Decode(v)
			},
		},
	},
	{
		"codec-msgpack",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				var h codec.MsgpackHandle
				var buf bytes.Buffer
				err := codec.NewEncoder(&buf, &h).Encode(v)
				return buf.Bytes(), err
			},
			Unmarshal: func(data []byte, v interface{}) error {
				var h codec.MsgpackHandle
				return codec.NewDecoderBytes(data, &h).Decode(v)
			},
		},
	},
	{
		"codec-cbor",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				var h codec.CborHandle
				var buf bytes.Buffer
				err := codec.NewEncoder(&buf, &h).Encode(v)
				return buf.Bytes(), err
			},
			Unmarshal: func(data []byte, v interface{}) error {
				var h codec.CborHandle
				return codec.NewDecoderBytes(data, &h).Decode(v)
			},
		},
	},
	{
		"fxamacker-cbor",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				return cbor.Marshal(v, cbor.EncOptions{Canonical: false})
			},
			Unmarshal: func(data []byte, v interface{}) error {
				return cbor.Unmarshal(data, v)
			},
		},
	},
	{
		"fxamacker-cbor-v2",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				//return cborv2.Marshal(v, cborv2.EncOptions{})
				return cborv2Encoder.Marshal(v)
			},
			Unmarshal: func(data []byte, v interface{}) error {
				return cborv2.Unmarshal(data, v)
			},
		},
	},
	{
		"bson",
		&testMarshaller{
			Marshal:   bson.Marshal,
			Unmarshal: bson.Unmarshal,
		},
	},
	{
		"gogo-proto",
		&testMarshaller{
			Marshal: func(v interface{}) ([]byte, error) {
				return v.(proto.Marshaler).Marshal()
			},
			Unmarshal: func(data []byte, v interface{}) error {
				return v.(proto.Unmarshaler).Unmarshal(data)
			},
		},
	},
}
