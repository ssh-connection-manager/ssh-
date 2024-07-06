package view

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Table struct {
	HeaderStyle  lipgloss.Color
	OddCellStyle lipgloss.Color
	EvenRowStyle lipgloss.Color
	BorderStyle  lipgloss.Color
}

const (
	WidthTable   = 70
	TopPadding   = 0
	RightPadding = 1
	CellWidth    = 1
)

func (t *Table) ViewTable(value [][]string, headers []string) *table.Table {
	render := lipgloss.NewRenderer(os.Stdout)

	HeaderStyle := render.NewStyle().Foreground(t.HeaderStyle).Bold(true).Align(lipgloss.Center)
	CellStyle := render.NewStyle().Padding(TopPadding, RightPadding).Width(CellWidth)

	OddRowStyle := CellStyle.Foreground(t.OddCellStyle)
	EvenRowStyle := CellStyle.Foreground(t.EvenRowStyle)

	BorderStyle := lipgloss.NewStyle().Foreground(t.BorderStyle)

	viewTable := table.New().
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
		Headers(headers...).
		Rows(value...).
		Width(WidthTable)

	return viewTable
}
