package cmd

import (
	"github.com/spf13/cobra"
	"ssh+/cmd/connect"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Вывод подключенных ssh соединений",
	Long: `Данная команда выводит список ваших добавленных соедениний по ssh, 
	в случае его отсутвия попросит добавить подключения`,
	Run: func(cmd *cobra.Command, args []string) {
		connect.ConsoleConnect("test4")
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
