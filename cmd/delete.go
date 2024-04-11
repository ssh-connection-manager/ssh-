package cmd

import (
	"ssh+/app/output"
	del "ssh+/cmd/delete"

	"fmt"
	"os"
	"ssh+/app/file"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Вывод подключенных ssh соединений",
	Long: `Данная команда выводит список ваших добавленных соедениний по ssh, 
	в случае его отсутвия попросит добавить подключения`,
	Run: func(cmd *cobra.Command, args []string) {
		var connects file.Connections
		aliases := connects.GetConnectionsAlias()

		if len(aliases) == 0 {
			output.GetOutError("Подключений не найдено")
			os.Exit(1)
		}

		aliases = append(connects.GetConnectionsAlias(), "exit")

		sp := selection.New("Выбирите подключение для удаление", aliases)
		sp.PageSize = 5
		sp.FilterPlaceholder = ""

		choice, err := sp.RunPrompt()
		if err != nil {
			fmt.Printf("Error: %v\n", err)

			os.Exit(1)
		}

		if choice == "exit" {
			output.GetOutSuccess("Вы успешно вышли")
			os.Exit(1)
		}

		del.DeleteConnect(choice)
		fmt.Println("подключение удалено")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
