package cmd

import (
	"os"

	"ssh+/app/file"
	"ssh+/cmd/connect"
	"ssh+/view"

	"github.com/spf13/cobra"
)

var ConnectCmd = &cobra.Command{
	Use:   connect.Use,
	Short: connect.Short,
	Long:  connect.Long,
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
