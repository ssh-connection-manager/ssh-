package cmd

import (
	"fmt"

	"ssh+/app/output"
	"ssh+/cmd/list"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   list.UseCommand,
	Short: list.ShortDescription,
	Long:  list.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		list := list.Show()

		fmt.Println("A list of your connections:")

		for _, v := range list {
			output.GetOutSuccess(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
