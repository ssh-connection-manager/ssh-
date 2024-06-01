package json

import (
	"encoding/json"
	"os"

	"ssh+/app/file"
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
	filePath, err := file.GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		output.GetOutError("Ошибка получения путя к файлу")
	}

	c.SerializationJson(file.ReadFile(filePath))

	encodedConnect := SetCryptData(connect)

	c.Connects = append(c.Connects, encodedConnect)

	file.WriteFile(getPathToConnectFile(), c.deserializationJson())
}

func (c *Connections) DeleteConnectToJson(alias string) {
	filePath, err := file.GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		output.GetOutError("Ошибка получения путя к файлу")
	}

	c.SerializationJson(file.ReadFile(filePath))

	for i, v := range c.Connects {
		if v.Alias == alias {
			c.deleteJsonDataByIndex(i)

			file.WriteFile(getPathToConnectFile(), c.deserializationJson())

			return
		}
	}

	output.GetOutError("Не найдено подключение")
}
