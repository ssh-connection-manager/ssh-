package file

import (
	"os"
)

func (c *Connections) GetConnectionsAlias() []string {
	c.SerializationJson(ReadFile())

	var result []string

	for _, conn := range c.Connects {
		result = append(result, conn.Alias)
	}

	return result
}

func (c *Connections) deleteJsonDataByIndex(index int) {
	copy(c.Connects[index:], c.Connects[index+1:])

	c.Connects[len(c.Connects)-1] = Connect{}
	c.Connects = c.Connects[:len(c.Connects)-1]
}

func getPathToFile() string {
	return os.Getenv("PATH_FILE") +
		os.Getenv("SEPARATOR") +
		os.Getenv("NAME_FILE")
}

func getFullPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := homePath + getPathToFile()
	return path, nil
}
