package main

import (
	"ssh+/cmd"
	"ssh+/inits"
)

func init() {
	inits.LoadEnv()
	inits.GenerateCryptKey()
	inits.CreateFileConnects()
	inits.GenerateConfigFile()
}

func main() {
	cmd.Execute()
}
