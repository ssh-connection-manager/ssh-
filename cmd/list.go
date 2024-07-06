package cmd

import (
	"fmt"

	"ssh+/cmd/list"
	"ssh+/view"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   list.UseCommand,
	Short: list.ShortDescription,
	Long:  list.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		list := list.Show()

		fmt.Println("A list of your connections:")

		lightGray := lipgloss.Color("241")
		purple := lipgloss.Color("206")
		gray := lipgloss.Color("245")

		headers := []string{
			"Alias",
			"Updated at",
			"Created at",
		}

		table := view.Table{
			HeaderStyle:  purple,
			OddCellStyle: gray,
			EvenRowStyle: lightGray,
			BorderStyle:  purple,
		}

		fmt.Println(table.ViewTable(list, headers))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
