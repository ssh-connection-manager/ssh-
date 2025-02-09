package cmd

import (
	"ssh+/app/view"

	delCmd "ssh+/cmd/delete"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/kernel/v2/pkg/app"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
)

var deleteCmd = &cobra.Command{
	Use:   delCmd.UseCommand,
	Short: delCmd.ShortDescription,
	Long:  delCmd.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var connects json.Connections

		aliases, err := connects.GetConnectionsAlias()
		if err != nil {
			output.GetOutError(err.Error())
		}

		customChoice := view.Select{
			FilterPlaceholder: delCmd.FilterPlaceholder,
			SelectionPrompt:   delCmd.SelectionPrompt,
			FilterPrompt:      delCmd.FilterPrompt,
			Template:          delCmd.Template,
			PageSize:          delCmd.PageSize,
		}

		choice, err := customChoice.SelectedValue(aliases)
		if err != nil {
			output.GetOutError("Error selecting connection at delete: " + err.Error())
		}

		app.Delete(choice)

		output.GetOutSuccess("Connection removed")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
