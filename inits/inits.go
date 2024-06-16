package inits

import (
	"ssh+/app/crypt"
	"ssh+/app/json"
	"ssh+/config"
)

func generateConfigFile() {
	config.Generate()
}

func createFileConnects() {
	json.Generate()
}

func generateCryptKey() {
	crypt.GenerateKey()
}

func SetDependencies() {
	generateConfigFile()
	createFileConnects()
	generateCryptKey()
}
