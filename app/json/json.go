package json

import (
	"encoding/json"
	"ssh+/app/output"
)

func (c *Connections) SerializationJson(dataConnectsInFile string) {
	err := json.Unmarshal([]byte(dataConnectsInFile), &c)

	if err != nil {
		output.GetOutError("Error of json serialization")
	}
}

func (c *Connections) deserializationJson() []byte {
	newDataConnect, err := json.MarshalIndent(&c, "", " ")

	if err != nil {
		output.GetOutError("Error of json deserialization")
	}

	return newDataConnect
}
