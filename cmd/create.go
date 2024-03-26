package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ssh+/cmd/create"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Создет подключение",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		create.CreateConnect("test5", "test4", "test", "rest")

		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
