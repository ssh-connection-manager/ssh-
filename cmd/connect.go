package cmd

import (
	"fmt"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"os"
	"ssh+/app/file"
	"ssh+/cmd/connect"
	"strings"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Вывод подключенных ssh соединений",
	Long: `Данная команда выводит список ваших добавленных соедениний по ssh, 
	в случае его отсутвия попросит добавить подключения`,
	Run: func(cmd *cobra.Command, args []string) {
		var connects file.Connections
		aliases := connects.GetConnectionsAlias()

		const (
			customTemplate = `
{{- if .Prompt -}}
  {{ Bold .Prompt }}
{{ end -}}

{{ if .IsFiltered }}
  {{- print .FilterPrompt " " .FilterInput }}
{{ end }}

{{- range  $i, $choice := .Choices }}
  {{- if IsScrollUpHintPosition $i }}
    {{- print "⇡ " -}}
  {{- else if IsScrollDownHintPosition $i -}}
    {{- print "⇣ " -}} 
  {{- else -}}
    {{- print "  " -}}
  {{- end -}} 

  {{- if eq $.SelectedIndex $i }}
   {{- print (Foreground "201" (Bold "▸ ")) (Selected $choice) "\n" }}
  {{- else }}
    {{- print " " (Unselected $choice) "\n" }}
  {{- end }}
{{- end}}`
		)

		sp := selection.New("Выбор подключения", aliases)
		sp.PageSize = 5
		sp.FilterPlaceholder = " введите название подключения"
		sp.FilterPrompt = "Поиск по алиуса:"
		sp.Template = customTemplate
		sp.Filter = func(filter string, choice *selection.Choice[string]) bool {
			return strings.HasPrefix(choice.Value, filter)
		}

		blue := termenv.String().Foreground(termenv.ANSI256Color(206))

		sp.SelectedChoiceStyle = func(c *selection.Choice[string]) string {
			return blue.Bold().Styled(c.Value)
		}

		choice, err := sp.RunPrompt()

		if err != nil {
			fmt.Printf("Error: %v\n", err)

			os.Exit(1)
		}

		connect.ConsoleConnect(choice)
	},
}

func init() {
	rootCmd.AddCommand(ConnectCmd)
}
