package init

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
}

func createFileConnects() {
	file.GenerateFile(os.Getenv("FILE_NAME_CONNECTS"))
	json.CreateBaseJsonDataToFile()
}

func SetDependencies() {
	loadEnv()
	generateCryptKey()
	generateConfigFile()
	createFileConnects()
}
