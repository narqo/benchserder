package benchserder

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleEvent_MarshalUnmarshal() {
	for _, tf := range testEvents {
		data, _ := readFile(tf)
		event := &Event{}
		err := json.Unmarshal(data, event)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(os.Stdout, "source %s\n", tf)
		fmt.Fprintf(os.Stdout, "--- raw:\t\t\t %d\n", len(data))

		data, _ = json.Marshal(event)
		fmt.Fprintf(os.Stdout, "--- json:\t\t\t %d\n", len(data))

		data, _ = marshalEasyJSON(event)
		fmt.Fprintf(os.Stdout, "--- easyjson:\t\t\t %d\n", len(data))

		data, _ = marshalMsgpack(event)
		fmt.Fprintf(os.Stdout, "--- vmihailenco/msgpack:\t %d\n", len(data))

		data, _ = marshalMsgpackAsArray(event)
		fmt.Fprintf(os.Stdout, "--- vmihailenco/msgp (as array): %d\n", len(data))

		data, _ = marshalCodecMsgpack(event)
		fmt.Fprintf(os.Stdout, "--- codec/msgpack:\t\t %d\n", len(data))

		data, _ = marshalCodecCBOR(event)
		fmt.Fprintf(os.Stdout, "--- codec/cbor:\t\t\t %d\n", len(data))

		data, _ = marshalFxamackerCBOR(event)
		fmt.Fprintf(os.Stdout, "--- fxamacker/cbor:\t\t %d\n", len(data))

		data, _ = marshalBSON(event)
		fmt.Fprintf(os.Stdout, "--- bson:\t\t\t %d\n", len(data))

		data, _ = marshalGogoProto(event)
		fmt.Fprintf(os.Stdout, "--- gogoproto:\t\t\t %d\n", len(data))

		// Output:
	}
}
