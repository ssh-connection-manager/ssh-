package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"os"

	"ssh+/app/file"
	"ssh+/app/output"
)

func GenerateKey() {
	file.GenerateFile(os.Getenv("FILE_NAME_CRYPT_KEY"))

	filePath, err := file.GetFullPath(os.Getenv("FILE_NAME_CRYPT_KEY"))
	if err != nil {
		output.GetOutError("Ошибка получения файла rsa ключа")
	}

	data := file.ReadFile(filePath)

	if len(data) == 0 {
		key, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			output.GetOutError("Ошибка создания ключа")
		}

		keyBytes := x509.MarshalPKCS1PrivateKey(key)
		if err != nil {
			output.GetOutError("Ошибка преобразования ключа")
		}

		file.WriteFile(filePath, keyBytes)
	}
}
