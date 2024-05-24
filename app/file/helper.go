package file

import (
	"os"

	"ssh+/app/output"
)

func (c *Connections) GetConnectionsAlias() []string {
	filePath, err := GetFullPath(os.Getenv("FILE_NAME_CONNECTS"))
	if err != nil {
		output.GetOutError("Ошибка получения путя к файлу")
	}

	c.SerializationJson(ReadFile(filePath))

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

func getPathToFile() string {
	return os.Getenv("PATH_FILE") + os.Getenv("SEPARATOR")
}

func GetFullPath(fileName string) (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := homePath + getPathToFile() + fileName
	return path, nil
}

func GenerateFile(filename string) {
	filePath, err := GetFullPath(filename)
	if err != nil {
		output.GetOutError("Ошибка получения файла")
	}

	CreateFile(filePath)
}
