package json

import (
	"ssh+/app/file"

	"github.com/spf13/viper"
)

func Generate() {
	fileName := viper.GetString("NameFileConnects")

	if !file.IsExistFile(fileName) {
		file.GenerateFile(fileName)
		CreateBaseJsonDataToFile()
	}
}
