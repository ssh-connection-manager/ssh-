package cmd

import (
	"github.com/spf13/viper"
	"ssh+/app/output"
	"ssh+/cmd/change"
	"ssh+/view"

	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/json"
)

var changeCmd = &cobra.Command{
	Use:   change.UseCommand,
	Short: change.ShortDescription,
	Long:  change.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var alias, address, login, password string
		var connects json.Connections

		filePath := viper.GetString("FullPathConfig")
		fileName := viper.GetString("NameFileConnects")
		fileKey := viper.GetString("NameFileCryptKey")

		aliases, err := connects.GetConnectionsAlias(filePath, fileName, fileKey)
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

		choice := customChoice.SelectedValue(aliases)

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

		customTextInput.DrawInput()

		change.Connect(choice, alias, address, login, password)

		output.GetOutSuccess("Update called")
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}
