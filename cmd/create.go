package cmd

import (
	"fmt"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
	"os"
	"ssh+/cmd/create"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Создет подключение",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var alias, address, login, password string
		var input *textinput.TextInput
		var err error

		arguments := map[string]*string{
			"Алиас":  &alias,
			"Адресс": &address,
			"Логин":  &login,
			"Пароль": &password,
		}

		for name, arg := range arguments {
			input = textinput.New(name)
			input.Placeholder = name + "не должен быть пустым"

			*arg, err = input.RunPrompt()

			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		}

		create.CreateConnect(alias, address, login, password)

		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
