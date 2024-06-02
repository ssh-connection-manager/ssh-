package cmd

import (
	"ssh+/app/json"
	"ssh+/app/output"
	"ssh+/view"

	del "ssh+/cmd/delete"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   del.UseCommand,
	Short: del.ShortDescription,
	Long:  del.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var connects json.Connections

		aliases := connects.GetConnectionsAlias()

		customChoice := view.Select{
			FilterPlaceholder: del.FilterPlaceholder,
			SelectionPrompt:   del.SelectionPrompt,
			FilterPrompt:      del.FilterPrompt,
			Template:          del.Template,
			PageSize:          del.PageSize,
		}

		choice := customChoice.SelectedValue(aliases)

		del.Connect(choice)

		output.GetOutSuccess("подключение удалено")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
