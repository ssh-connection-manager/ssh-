package main

import (
	"ssh+/cmd"

	"github.com/ssh-connection-manager/kernel/v2/pkg/inits"
)

func init() {
	inits.SetDependencies()
}

func main() {
	cmd.Execute()
}
