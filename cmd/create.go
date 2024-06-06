package cmd

import (
	"ssh+/app/output"
	"ssh+/cmd/create"
	"ssh+/view"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   create.UseCommand,
	Short: create.ShortDescription,
	Long:  create.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var alias, address, login, password string

		var arguments = [][]*string{
			{&create.NameAlias, &alias},
			{&create.NameAddress, &address},
			{&create.NameLogin, &login},
			{&create.NamePassword, &password},
		}

		hiddenArgs := []*string{&password}

		customTextInput := view.TextInput{
			Placeholder: create.Placeholder,
			HiddenArgs:  hiddenArgs,
			Arguments:   arguments,
		}

		customTextInput.DrawInput()

		create.Connect(alias, address, login, password)

		output.GetOutSuccess("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
