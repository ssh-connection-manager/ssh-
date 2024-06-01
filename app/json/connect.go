package json

import (
	"os"
	"ssh+/app/file"
	"ssh+/app/output"
)

func (c *Connections) GetConnectionsAlias() []string {
	filePath, err := file.GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		output.GetOutError("Ошибка получения путя к файлу")
	}

	c.SerializationJson(file.ReadFile(filePath))

	var result []string

	for _, conn := range c.Connects {
		result = append(result, conn.Alias)
	}

	if len(result) == 0 {
		output.GetOutError("Подключений не найдено")
	}

	return result
}

func (c *Connections) deleteJsonDataByIndex(index int) {
	copy(c.Connects[index:], c.Connects[index+1:])

	c.Connects[len(c.Connects)-1] = Connect{}
	c.Connects = c.Connects[:len(c.Connects)-1]
}

func getPathToConnectFile() string {
	fullPath, err := file.GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		panic(err)
	}

	return fullPath
}
