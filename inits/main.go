package inits

import (
	"os"
	"ssh+/app/crypt"
	"ssh+/app/file"
	"ssh+/app/json"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GenerateCryptKey() {
	crypt.GenerateKey()
}

func CreateFileConnects() {
	file.GenerateFile(os.Getenv("FILE_NAME_CONNECTS"))
	json.CreateBaseJsonDataToFile()
}
