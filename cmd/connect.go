package cmd

import (
	"fmt"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/spf13/cobra"
	"os"
	"ssh+/app/file"
	"ssh+/cmd/connect"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Вывод подключенных ssh соединений",
	Long: `Данная команда выводит список ваших добавленных соедениний по ssh, 
	в случае его отсутвия попросит добавить подключения`,
	Run: func(cmd *cobra.Command, args []string) {
		var connects file.Connections
		aliases := connects.GetConnectionsAlias()

		sp := selection.New("Выбирите подключение для удаление", aliases)
		sp.PageSize = 5
		sp.FilterPlaceholder = ""

		choice, err := sp.RunPrompt()

		if err != nil {
			fmt.Printf("Error: %v\n", err)

			os.Exit(1)
		}

		connect.ConsoleConnect(choice)
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
