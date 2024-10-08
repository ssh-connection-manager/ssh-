package cmd

import (
	"ssh+/app/output"
	"ssh+/view"

	crCmd "ssh+/cmd/create"

	"github.com/spf13/cobra"
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

		customTextInput.DrawInput()

		crCmd.Connect(alias, address, login, password)

		output.GetOutSuccess("Create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
