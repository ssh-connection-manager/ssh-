package json

import (
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

func getPathToConnectFile() string {
	fullPath, err := file.GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		panic(err)
	}

	return fullPath
}

func (c *Connections) GetConnectionsAlias() []string {
	filePath := getPathToConnectFile()

	c.SerializationJson(file.ReadFile(filePath))
	c.SetDecryptData()

	var result []string

	for _, conn := range c.Connects {
		result = append(result, conn.Alias)
	}

	if len(result) == 0 {
		output.GetOutError("Подключений не найдено")
	}

	return result
}

func (c *Connections) WriteConnectToJson(connect Connect) {
	filePath := getPathToConnectFile()
	c.SerializationJson(file.ReadFile(filePath))

	encodedConnect := SetCryptData(connect)
	c.Connects = append(c.Connects, encodedConnect)

	file.WriteFile(getPathToConnectFile(), c.deserializationJson())
}

func (c *Connections) deleteJsonDataByIndex(index int) {
	copy(c.Connects[index:], c.Connects[index+1:])

	c.Connects[len(c.Connects)-1] = Connect{}
	c.Connects = c.Connects[:len(c.Connects)-1]
}

func (c *Connections) DeleteConnectToJson(alias string) {
	filePath := getPathToConnectFile()

	c.SerializationJson(file.ReadFile(filePath))
	c.SetDecryptData()

	for i, v := range c.Connects {
		if v.Alias == alias {
			c.deleteJsonDataByIndex(i)

			file.WriteFile(getPathToConnectFile(), c.deserializationJson())

			return
		}
	}

	output.GetOutError("Не найдено подключение")
}
