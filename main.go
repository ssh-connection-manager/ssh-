package main

import (
	"ssh+/cmd"
	"ssh+/inits"
)

func init() {
	inits.LoadEnv()
	inits.GenerateCryptKey()
	inits.CreateFileConnects()
}

func main() {
	cmd.Execute()
}
