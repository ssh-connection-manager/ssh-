package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"ssh+/cmd/root"
)

var rootCmd = &cobra.Command{
	Use:   root.UseCommand,
	Short: root.ShortDescription,
	Long:  root.LongDescription,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
