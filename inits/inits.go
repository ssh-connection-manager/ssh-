package inits

import (
	"os"

	"ssh+/app/crypt"
	"ssh+/app/file"
	"ssh+/app/json"
	"ssh+/config"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func generateCryptKey() {
	crypt.GenerateKey()
}

func generateConfigFile() {
	config.Generate()
	config.SetConfigVariable()

}

func createFileConnects() {
	file.GenerateFile(os.Getenv("FILE_NAME_CONNECTS"))
	json.CreateBaseJsonDataToFile()
}

func SetDependencies() {
	generateConfigFile()
	loadEnv()
	generateCryptKey()
	createFileConnects()
}
