package inits

import (
	"github.com/joho/godotenv"
	"os"
	"ssh+/app/encrypt"
	"ssh+/app/file"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GenerateCryptKey() {
	encrypt.GenerateKey()
}

func CreateFileConnects() {
	file.GenerateFile(os.Getenv("FILE_NAME_CONNECTS"))
}
