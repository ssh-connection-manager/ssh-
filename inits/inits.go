package inits

import (
	"github.com/spf13/viper"
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
	file.GenerateFile(viper.GetString("NameFileConnects"))
	json.CreateBaseJsonDataToFile()
}

func SetDependencies() {
	generateConfigFile()
	loadEnv()
	generateCryptKey()
	createFileConnects()
}
