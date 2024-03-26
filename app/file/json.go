package file

import (
	"encoding/json"
	"fmt"
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

func (c *Connections) ReadJsonData(jsonData string) {
	err := json.Unmarshal([]byte(jsonData), &c)
	if err != nil {
		output.GetOutError("пустой файл")
	}
}

func (c *Connections) WriteConnectToJson(connect Connect) {
	dataConnectsInFile := ReadFile()
	c.ReadJsonData(dataConnectsInFile)

	c.Connects = append(c.Connects, connect)

	newDataConnect, err := json.MarshalIndent(&c, "", " ")

	if err != nil {
		output.GetOutError("Ошибка записи json")
	}

	WriteFile(newDataConnect)
}

func (c *Connections) DeleteConnectToJson() {
	dataConnectsInFile := ReadFile()
	c.ReadJsonData(dataConnectsInFile)

	fmt.Println(c)
}
