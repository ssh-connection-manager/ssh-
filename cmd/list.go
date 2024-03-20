package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ssh+/app/list"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Вывод подключенных ssh соединений",
	Long: `Данная команда выводит список ваших добавленных соедениний по ssh, 
	в случае его отсутвия попросит добавить подключения`,
	Run: func(cmd *cobra.Command, args []string) {
		list := list.GetConnectList()
		fmt.Println(list)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
