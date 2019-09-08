package benchserder

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/fxamacker/cbor"
	proto "github.com/golang/protobuf/proto"
	"github.com/mailru/easyjson"
	"github.com/ugorji/go/codec"
	"github.com/vmihailenco/msgpack/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type marshalFunc func(interface{}) ([]byte, error)

type unmarshalFunc func([]byte, interface{}) error

func marshalEasyJSON(v interface{}) ([]byte, error) {
	return easyjson.Marshal(v.(easyjson.Marshaler))
}

func unmarshalEasyJSON(data []byte, v interface{}) error {
	return easyjson.Unmarshal(data, v.(easyjson.Unmarshaler))
}

func marshalMsgpack(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.UseJSONTag(true)
	err := enc.Encode(v)
	return buf.Bytes(), err
}

func marshalMsgpackAsArray(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.StructAsArray(true)
	err := enc.Encode(v)
	return buf.Bytes(), err
}

func unmarshalMsgpack(data []byte, v interface{}) error {
	dec := msgpack.NewDecoder(bytes.NewReader(data))
	dec.UseJSONTag(true)
	return dec.Decode(v)
}

func marshalCodecMsgpack(v interface{}) ([]byte, error) {
	var h codec.MsgpackHandle
	var buf bytes.Buffer
	enc := codec.NewEncoder(&buf, &h)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func unmarshalCodecMsgpack(data []byte, v interface{}) error {
	var h codec.MsgpackHandle
	dec := codec.NewDecoderBytes(data, &h)
	return dec.Decode(v)
}

func marshalCodecCBOR(v interface{}) ([]byte, error) {
	var h codec.CborHandle
	var buf bytes.Buffer
	enc := codec.NewEncoder(&buf, &h)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func unmarshalCodecCBOR(data []byte, v interface{}) error {
	var h codec.CborHandle
	dec := codec.NewDecoderBytes(data, &h)
	return dec.Decode(v)
}

func marshalFxamackerCBOR(v interface{}) ([]byte, error) {
	return cbor.Marshal(v, cbor.EncOptions{Canonical: false})
}

func unmarshalFxamackerCBOR(data []byte, v interface{}) error {
	return cbor.Unmarshal(data, v)
}

func marshalBSON(v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

func unmarshalBSON(data []byte, v interface{}) error {
	return bson.Unmarshal(data, v)
}

func marshalGogoProto(v interface{}) ([]byte, error) {
	return v.(proto.Marshaler).Marshal()
}

func unmarshalGogoProto(data []byte, v interface{}) error {
	return v.(proto.Unmarshaler).Unmarshal(data)
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
