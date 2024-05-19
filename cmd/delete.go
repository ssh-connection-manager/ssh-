package cmd

import (
	"os"

	"ssh+/app/file"
	"ssh+/app/output"
	"ssh+/view"

	del "ssh+/cmd/delete"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   del.Use,
	Short: del.Short,
	Long:  del.Long,
	Run: func(cmd *cobra.Command, args []string) {
		var connects file.Connections

		aliases := connects.GetConnectionsAlias()

		customChoice := view.Select{
			FilterPlaceholder: os.Getenv("DELETE_FILTER_PLACEHOLDER"),
			SelectionPrompt:   os.Getenv("DELETE_SELECTION_PROMPT"),
			FilterPrompt:      os.Getenv("DELETE_FILTER_PROMPT"),
			Template:          os.Getenv("DELETE_SELECT_TEMPLATE"),
			PageSize:          os.Getenv("DELETE_PAGE_SIZE"),
		}

		choice := customChoice.SelectedValue(aliases)

		del.DeleteConnect(choice)

		output.GetOutSuccess("подключение удалено")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
