package file

import (
	"github.com/spf13/viper"
	"os"

	"ssh+/app/output"
)

func getPathToFile() string {
	return viper.GetString("FullPathConfig") + viper.GetString("Separator")
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
