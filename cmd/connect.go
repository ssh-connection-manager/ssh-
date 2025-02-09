package cmd

import (
	"ssh+/app/view"
	"ssh+/cmd/connect"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/kernel/v2/pkg/app"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
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

		app.Connect(choice)
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
