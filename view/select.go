package view

import (
	"github.com/erikgeiser/promptkit/selection"
	"github.com/muesli/termenv"
	"os"
	"ssh+/app/output"
	"strconv"
	"strings"
)

type Select struct {
	FilterPlaceholder string
	SelectionPrompt   string
	FilterPrompt      string
	Template          string
	PageSize          string
}

func (s Select) SelectedValue(aliases []string) string {
	sp := selection.New(s.SelectionPrompt, aliases)

	sp.FilterPlaceholder = os.Getenv("SPACE") + s.FilterPlaceholder
	sp.FilterPrompt = s.FilterPrompt
	sp.Template = s.Template

	pageSize, err := strconv.Atoi(s.PageSize)

	if err != nil {
		output.GetOutError("Неправильная ковертация")
	}

	sp.PageSize = pageSize

	sp.Filter = func(filter string, choice *selection.Choice[string]) bool {
		return strings.HasPrefix(choice.Value, filter)
	}

	color := termenv.String().Foreground(termenv.ANSI256Color(206))

	sp.SelectedChoiceStyle = func(c *selection.Choice[string]) string {
		return color.Bold().Styled(c.Value)
	}

	selectedValue, err := sp.RunPrompt()

	if err != nil {
		output.GetOutError("Ошибка ввода")
	}

	return selectedValue
}
