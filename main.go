package main

import (
	"ssh+/cmd"

	"github.com/ssh-connection-manager/kernel/inits"
)

func init() {
	inits.SetDependencies()
}

func main() {
	cmd.Execute()
}
