package main

import (
	"ssh+/cmd"
	"ssh+/inits"
)

func init() {
	inits.SetDependencies()
}

func main() {
	cmd.Execute()
}
