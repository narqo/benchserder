package benchserder

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleEvent_MarshalUnmarshal() {
	for _, tf := range testEvents {
		jsonData, _ := readFile(tf)
		event := &Event{}
		err := json.Unmarshal(jsonData, event)
		if err != nil {
			panic(err)
		}

		m1Data, _ := marshalMsgpack(event)
		m2Data, _ := marshalCodecMsgpack(event)
		c1Data, _ := marshalCodecCBOR(event)
		c2Data, _ := marshalFxamackerCBOR(event)
		bData, _ := marshalBSON(event)
		pbData, _ := marshalGogoProto(event)

		fmt.Fprintf(os.Stdout, "source %s\n", tf)
		fmt.Fprintf(os.Stdout, "--- json:\t\t %d\n", len(jsonData))
		fmt.Fprintf(os.Stdout, "--- vmihailenco/msgpack: %d\n", len(m1Data))
		fmt.Fprintf(os.Stdout, "--- codec/msgpack:\t %d\n", len(m2Data))
		fmt.Fprintf(os.Stdout, "--- codec/cbor:\t\t %d\n", len(c1Data))
		fmt.Fprintf(os.Stdout, "--- fxamacker/cbor:\t %d\n", len(c2Data))
		fmt.Fprintf(os.Stdout, "--- bson:\t\t %d\n", len(bData))
		fmt.Fprintf(os.Stdout, "--- gogoproto:\t\t %d\n", len(pbData))

		// Output:
	}
}
