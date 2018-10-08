package jsonutils

import (
	"encoding/json"
	"fmt"
)

// stringMapToJSON takes a map[string]string and turns it into a *[]byte element in JSON format.
func StringMapToJSON(data map[string]string, name string) *[]byte {
	var jsonCaps []byte
	var err error

	if name != "" {
		mapWithName := map[string]map[string]string{
			name: data,
		}
		jsonCaps, err = json.Marshal(mapWithName)
	} else {
		if data == nil {
			jsonCaps, err = []byte("{}"), nil
		} else {
			jsonCaps, err = json.Marshal(data)
		}
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("JSON Sent:\t" + string(jsonCaps))

	return &jsonCaps
}

// jsonToMap takes a *[]byte and turns it into a map[string]string element.
func JSONToMap(body *[]byte) map[string]*json.RawMessage {
	fmt.Println("JSON Recieved:\t" + string(*body))

	var result map[string]*json.RawMessage

	err := json.Unmarshal(*body, &result)
	if err != nil {
		panic(err)
	}

	return result
}
