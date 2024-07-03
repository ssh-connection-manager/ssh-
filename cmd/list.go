package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"os"

	"ssh+/cmd/list"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   list.UseCommand,
	Short: list.ShortDescription,
	Long:  list.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		list := list.Show()

		fmt.Println("A list of your connections:")

		re := lipgloss.NewRenderer(os.Stdout)

		purple := lipgloss.Color("99")
		gray := lipgloss.Color("245")
		lightGray := lipgloss.Color("241")

		// HeaderStyle is the lipgloss style used for the table headers.
		HeaderStyle := re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		// CellStyle is the base lipgloss style used for the table rows.
		CellStyle := re.NewStyle().Padding(0, 1).Width(14)
		// OddRowStyle is the lipgloss style used for odd-numbered table rows.
		OddRowStyle := CellStyle.Foreground(gray)
		// EvenRowStyle is the lipgloss style used for even-numbered table rows.
		EvenRowStyle := CellStyle.Foreground(lightGray)
		// BorderStyle is the lipgloss style used for the table border.
		BorderStyle := lipgloss.NewStyle().Foreground(purple)

		var list2 [][]string

		for _, v := range list {
			newElement := []string{v}
			list2 = append(list2, newElement)
		}

		fmt.Println(list)

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(BorderStyle).
			StyleFunc(func(row, col int) lipgloss.Style {
				var style lipgloss.Style

				switch {
				case row == 0:
					return HeaderStyle
				case row%2 == 0:
					style = EvenRowStyle
				default:
					style = OddRowStyle
				}

				return style
			}).
			Headers("Alias", "Updated at", "Created at").
			Rows(list2...).Width(50)

		fmt.Println(t)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
