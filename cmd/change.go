package cmd

import (
	"ssh+/cmd/change"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/output"
	"github.com/ssh-connection-manager/view"
)

var changeCmd = &cobra.Command{
	Use:   change.UseCommand,
	Short: change.ShortDescription,
	Long:  change.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var alias, address, login, password string
		var connects json.Connections

		aliases, err := connects.GetConnectionsAlias()
		if err != nil {
			output.GetOutError(err.Error())
		}

		customChoice := view.Select{
			FilterPlaceholder: change.FilterPlaceholder,
			SelectionPrompt:   change.SelectionPrompt,
			FilterPrompt:      change.FilterPrompt,
			Template:          change.Template,
			PageSize:          change.PageSize,
		}

		choice, err := customChoice.SelectedValue(aliases)
		if err != nil {
			output.GetOutError("Error selecting alias: " + err.Error())
		}

		change.ExistByIndex(choice)

		arguments := [][]*string{
			{&change.NameAlias, &alias},
			{&change.NameAddress, &address},
			{&change.NameLogin, &login},
			{&change.NamePassword, &password},
		}

		hiddenArgs := []*string{&password}

		customTextInput := view.TextInput{
			Placeholder: change.Placeholder,
			HiddenArgs:  hiddenArgs,
			Arguments:   arguments,
		}

		err = customTextInput.DrawInput()
		if err != nil {
			output.GetOutError("Error drawing input at change: " + err.Error())
		}

		change.Connect(choice, alias, address, login, password)

		output.GetOutSuccess("Update called")
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}
