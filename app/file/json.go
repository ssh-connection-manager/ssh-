package file

import (
	"encoding/json"

	"ssh+/app/output"
)

type Connections struct {
	Connects []Connect `json:"connects"`
}

type Connect struct {
	Alias    string `json:"alias"`
	Login    string `json:"login"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (c *Connections) SerializationJson(dataConnectsInFile string) {
	err := json.Unmarshal([]byte(dataConnectsInFile), &c)

	if err != nil {
		output.GetOutError("Ошибка сериализации json")
	}
}

func (c *Connections) deserializationJson() []byte {
	newDataConnect, err := json.MarshalIndent(&c, "", " ")

	if err != nil {
		output.GetOutError("Ошибка десиарилизации json")
	}

	return newDataConnect
}

func (c *Connections) WriteConnectToJson(connect Connect) {
	c.SerializationJson(ReadFile())

	c.Connects = append(c.Connects, connect)

	WriteFile(c.deserializationJson())
}

func (c *Connections) DeleteConnectToJson(alias string) {
	c.SerializationJson(ReadFile())

	for i, v := range c.Connects {
		if v.Alias == alias {
			c.deleteJsonDataByIndex(i)

			WriteFile(c.deserializationJson())

			return
		}
	}

	output.GetOutError("Не найдено подключение")
}
