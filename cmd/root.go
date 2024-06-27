package cmd

import (
	"ssh+/app/output"
	"ssh+/cmd/root"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: root.UseCommand,
}

func Execute() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	rootCmd.Long = root.GetLongDescription()

	err := rootCmd.Execute()
	if err != nil {
		output.GetOutError("ssh+ is die")
	}
}
