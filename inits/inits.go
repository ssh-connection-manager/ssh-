package inits

import (
	"ssh+/app/output"
	"ssh+/config"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/crypt"
	"github.com/ssh-connection-manager/json"
)

func generateConfigFile() {
	config.Generate()
}

func createFileConnects() {
	pathConf := viper.GetString("FullPathConfig")
	confName := viper.GetString("NameFileConnects")

	err := json.Generate(pathConf, confName)
	if err != nil {
		output.GetOutError("err create file connect")
	}
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
