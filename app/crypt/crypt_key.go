package crypt

import (
	"crypto/rand"
	"ssh+/app/output"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/file"
)

func getPathToKey() string {
	filePath := file.GetFullPath(
		viper.GetString("FullPathConfig"),
		viper.GetString("NameFileCryptKey"))

	return filePath
}

func GetKey() []byte {
	data, err := file.ReadFile(getPathToKey())
	if err != nil {
		output.GetOutError("File opening error")
	}

	return []byte(data)
}

func GenerateKey() {
	filePath := file.GetFullPath(
		viper.GetString("FullPathConfig"),
		viper.GetString("NameFileCryptKey"))

	if !file.IsExistFile(filePath) {
		file.CreateFile(filePath)

		if len(GetKey()) == 0 {
			key := make([]byte, 32)

			_, err := rand.Read(key)
			if err != nil {
				output.GetOutError("Key generation error")
			}

			err = file.WriteFile(getPathToKey(), key)
			if err != nil {
				output.GetOutError("Error writing key")
			}
		}
	}
}
