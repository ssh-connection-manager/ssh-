package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"ssh+/app/file"
	"ssh+/cmd/connect"
	"ssh+/view"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Вывод подключенных ssh соединений",
	Long: `Данная команда выводит список ваших добавленных соедениний по ssh, 
	в случае его отсутвия попросит добавить подключения`,
	Run: func(cmd *cobra.Command, args []string) {
		var connects file.Connections

		aliases := connects.GetConnectionsAlias()

		customChoice := view.Select{
			FilterPlaceholder: os.Getenv("CONNECT_FILTER_PLACEHOLDER"),
			SelectionPrompt:   os.Getenv("CONNECT_SELECTION_PROMPT"),
			FilterPrompt:      os.Getenv("CONNECT_FILTER_PROMPT"),
			Template:          os.Getenv("CONNECT_SELECT_TEMPLATE"),
			PageSize:          os.Getenv("CONNECT_PAGE_SIZE"),
		}

		choice := customChoice.SelectedValue(aliases)

		connect.ConsoleConnect(choice)
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
