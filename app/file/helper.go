package file

import (
	"os"

	"ssh+/app/output"
)

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
