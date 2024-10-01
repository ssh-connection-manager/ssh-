package inits

import (
	"github.com/spf13/viper"
	"ssh+/app/json"
	"ssh+/app/output"
	"ssh+/config"

	"github.com/ssh-connection-manager/crypt"
)

func generateConfigFile() {
	config.Generate()
}

func createFileConnects() {
	json.Generate()
}

func generateCryptKey() {
	pathConf := viper.GetString("FullPathConfig")
	fileNameKey := viper.GetString("NameFileCryptKey")

	err := crypt.GenerateKey(pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err generate key")
	}
}

func SetDependencies() {
	generateConfigFile()
	createFileConnects()
	generateCryptKey()
}
