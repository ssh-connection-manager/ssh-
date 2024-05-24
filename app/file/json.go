package file

import (
	"encoding/json"
	"os"

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

func getPathToConnectFile() string {
	fullPath, err := GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		panic(err)
	}

	return fullPath
}

func (c *Connections) WriteConnectToJson(connect Connect) {
	filePath, err := GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		output.GetOutError("Ошибка получения путя к файлу")
	}

	c.SerializationJson(ReadFile(filePath))

	c.Connects = append(c.Connects, connect)

	WriteFile(getPathToConnectFile(), c.deserializationJson())
}

func (c *Connections) DeleteConnectToJson(alias string) {
	filePath, err := GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		output.GetOutError("Ошибка получения путя к файлу")
	}

	c.SerializationJson(ReadFile(filePath))

	for i, v := range c.Connects {
		if v.Alias == alias {
			c.deleteJsonDataByIndex(i)

			WriteFile(getPathToConnectFile(), c.deserializationJson())

			return
		}
	}

	output.GetOutError("Не найдено подключение")
}
