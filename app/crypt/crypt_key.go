package crypt

import (
	"crypto/rand"
	"ssh+/app/file"
	"ssh+/app/output"

	"github.com/spf13/viper"
)

func getPathToKey() string {
	filePath := file.GetFullPath(viper.GetString("NameFileCryptKey"))

	return filePath
}

func GetKey() []byte {
	data, err := file.ReadFile(getPathToKey())
	if err != nil {
		output.GetOutError("Ошибка открытия файла")
	}

	return []byte(data)
}

func GenerateKey() {
	fileName := viper.GetString("NameFileCryptKey")

	if !file.IsExistFile(fileName) {
		file.GenerateFile(fileName)

		if len(GetKey()) == 0 {
			key := make([]byte, 32)

			_, err := rand.Read(key)
			if err != nil {
				output.GetOutError("Ошибка генирации ключа")
			}

			file.WriteFile(getPathToKey(), key)
		}
	}
}
