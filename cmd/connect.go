package cmd

import (
	"ssh+/cmd/connect"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/output"
	"github.com/ssh-connection-manager/view"
)

var ConnectCmd = &cobra.Command{
	Use:   connect.UseCommand,
	Short: connect.ShortDescription,
	Long:  connect.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var connects json.Connections

		aliases, err := connects.GetConnectionsAlias()
		if err != nil {
			output.GetOutError(err.Error())
		}

		customChoice := view.Select{
			FilterPlaceholder: connect.FilterPlaceholder,
			SelectionPrompt:   connect.SelectionPrompt,
			FilterPrompt:      connect.FilterPrompt,
			Template:          connect.Template,
			PageSize:          connect.PageSize,
		}

		choice, err := customChoice.SelectedValue(aliases)
		if err != nil {
			output.GetOutError("Error selecting connection at connect: " + err.Error())
		}

		connect.Ssh(choice)
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
