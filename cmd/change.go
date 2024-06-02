package cmd

import (
	"ssh+/app/json"
	"ssh+/app/output"
	"ssh+/cmd/change"
	"ssh+/view"

	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   change.UseCommand,
	Short: change.ShortDescription,
	Long:  change.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var alias, address, login, password string
		var connects json.Connections

		aliases := connects.GetConnectionsAlias()

		customChoice := view.Select{
			FilterPlaceholder: change.FilterPlaceholder,
			SelectionPrompt:   change.SelectionPrompt,
			FilterPrompt:      change.FilterPrompt,
			Template:          change.Template,
			PageSize:          change.PageSize,
		}

		choice := customChoice.SelectedValue(aliases)

		change.ExistByIndex(choice)

		arguments := map[string]*string{
			"Алиас":  &alias,
			"Адресс": &address,
			"Логин":  &login,
			"Пароль": &password,
		}

		hiddenArgs := []*string{&password}

		customTextInput := view.TextInput{
			Placeholder: change.Placeholder,
			HiddenArgs:  hiddenArgs,
		}

		customTextInput.DrawInput(arguments)

		change.Connect(choice, alias, address, login, password)

		output.GetOutSuccess("update called")
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}
