package main

import (
	"ssh+/cmd"
	"ssh+/init"
)

func init() {
	init.SetDependencies()
}

func main() {
	cmd.Execute()
}
