package json

import (
	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/file"
)

func Generate() {
	filePath := file.GetFullPath(
		viper.GetString("FullPathConfig"),
		viper.GetString("NameFileConnects"))

	if !file.IsExistFile(filePath) {
		file.CreateFile(filePath)
		CreateBaseJsonDataToFile()
	}
}
