package jsonutil

import (
	"encoding/json"

	"github.com/narqo/benchserder/internal/reflectutil"
)

var nullStr = []byte("null")

func Marshal(value interface{}) ([]byte, error) {
	if value, ok := value.(json.Marshaler); ok {
		if reflectutil.IsNilInterface(value) {
			return nullStr, nil
		}
		return value.MarshalJSON()
	}
	return json.Marshal(value)
}

func Unmarshal(data []byte, value interface{}) error {
	if _, ok := value.(*json.RawMessage); !ok {
		if value, ok := value.(json.Unmarshaler); ok {
			return value.UnmarshalJSON(data)
		}
	}
	return json.Unmarshal(data, value)
}

func UnmarshalString(data string, value interface{}) error {
	return Unmarshal(strToBytes(data), value)
}
