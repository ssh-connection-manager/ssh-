package main

import (
	"ssh+/cmd"
	"ssh+/inits"
)

func init() {
	inits.CreateFileConnects()
}

func main() {
	cmd.Execute()
}
