package cmd

import (
	"ssh+/app/output"
	"ssh+/cmd/connect"
	"ssh+/view"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/json"
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

		choice := customChoice.SelectedValue(aliases)

		connect.Ssh(choice)
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
