package cmd

import (
	"ssh+/app/output"
	"ssh+/view"

	del "ssh+/cmd/delete"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/json"
)

var deleteCmd = &cobra.Command{
	Use:   del.UseCommand,
	Short: del.ShortDescription,
	Long:  del.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var connects json.Connections

		filePath := viper.GetString("FullPathConfig")
		fileName := viper.GetString("NameFileConnects")

		aliases, err := connects.GetConnectionsAlias(filePath, fileName)
		if err != nil {
			output.GetOutError(err.Error())
		}

		customChoice := view.Select{
			FilterPlaceholder: del.FilterPlaceholder,
			SelectionPrompt:   del.SelectionPrompt,
			FilterPrompt:      del.FilterPrompt,
			Template:          del.Template,
			PageSize:          del.PageSize,
		}

		choice := customChoice.SelectedValue(aliases)

		del.Connect(choice)

		output.GetOutSuccess("Connection removed")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
