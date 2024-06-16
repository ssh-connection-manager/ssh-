package file

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func GetFullPath(fileName string) string {
	return filepath.Join(viper.GetString("FullPathConfig"), fileName)
}

func IsExistFile(filename string) bool {
	filePath := GetFullPath(filename)

	_, err := ReadFile(filePath)
	if err != nil {
		return false
	}

	return true
}

func GenerateFile(filename string) {
	filePath := GetFullPath(filename)

	CreateFile(filePath)
}
