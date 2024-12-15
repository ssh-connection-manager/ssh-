package cmd

import (
	"ssh+/cmd/root"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
)

var rootCmd = &cobra.Command{
	Use:  root.UseCommand,
	Long: root.LongDescription,
}

func Execute() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		output.GetOutError("ssh+ error")
	}
}
