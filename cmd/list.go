package cmd

import (
	"fmt"

	lsCmd "ssh+/cmd/list"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/ssh-connection-manager/kernel/app"
	"github.com/ssh-connection-manager/view"
)

var listCmd = &cobra.Command{
	Use:   lsCmd.UseCommand,
	Short: lsCmd.ShortDescription,
	Long:  lsCmd.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		list := app.List()

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
