package cmd

import (
	"fmt"

	"ssh+/app/output"
	"ssh+/cmd/list"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   list.Use,
	Short: list.Short,
	Long:  list.Long,
	Run: func(cmd *cobra.Command, args []string) {
		list := list.GetConnectsList()

		fmt.Println("Список ваших подключений:")

		for _, v := range list {
			output.GetOutSuccess(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
