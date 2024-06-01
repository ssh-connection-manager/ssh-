package crypt

import (
	"crypto/rand"
	"os"

	"ssh+/app/file"
	"ssh+/app/output"
)

func getPathToKey() string {
	filePath, err := file.GetFullPath(os.Getenv("FILE_NAME_CRYPT_KEY"))
	if err != nil {
		output.GetOutError("Ошибка получения файла rsa ключа")
	}

	return filePath
}

func GetKey() []byte {
	return []byte(file.ReadFile(getPathToKey()))
}

func GenerateKey() {
	file.GenerateFile(os.Getenv("FILE_NAME_CRYPT_KEY"))

	if len(GetKey()) == 0 {
		key := make([]byte, 32)

		_, err := rand.Read(key)
		if err != nil {
			output.GetOutError("Ошибка генирации ключа")
		}

		file.WriteFile(getPathToKey(), key)
	}
}
