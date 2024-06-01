package cmd

import (
	"ssh+/app/json"
	"ssh+/cmd/connect"
	"ssh+/view"

	"github.com/spf13/cobra"
)

var ConnectCmd = &cobra.Command{
	Use:   connect.UseCommand,
	Short: connect.ShortDescription,
	Long:  connect.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var connects json.Connections

		aliases := connects.GetConnectionsAlias()

		customChoice := view.Select{
			FilterPlaceholder: connect.FilterPlaceholder,
			SelectionPrompt:   connect.SelectionPrompt,
			FilterPrompt:      connect.FilterPrompt,
			Template:          connect.Template,
			PageSize:          connect.PageSize,
		}

		choice := customChoice.SelectedValue(aliases)

		connect.ConsoleConnect(choice)
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
