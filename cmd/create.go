package cmd

import (
	crCmd "ssh+/cmd/create"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/output"
	"github.com/ssh-connection-manager/view"
)

var createCmd = &cobra.Command{
	Use:   crCmd.UseCommand,
	Short: crCmd.ShortDescription,
	Long:  crCmd.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var alias, address, login, password string

		var arguments = [][]*string{
			{&crCmd.NameAlias, &alias},
			{&crCmd.NameAddress, &address},
			{&crCmd.NameLogin, &login},
			{&crCmd.NamePassword, &password},
		}

		hiddenArgs := []*string{&password}

		customTextInput := view.TextInput{
			Placeholder: crCmd.Placeholder,
			HiddenArgs:  hiddenArgs,
			Arguments:   arguments,
		}

		err := customTextInput.DrawInput()
		if err != nil {
			output.GetOutError("Error drawing input at create: " + err.Error())
		}

		crCmd.Connect(alias, address, login, password)

		output.GetOutSuccess("Create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
