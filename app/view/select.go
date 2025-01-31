package view

import (
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
	"strconv"
	"strings"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/muesli/termenv"
)

type Select struct {
	FilterPlaceholder string
	SelectionPrompt   string
	FilterPrompt      string
	Template          string
	PageSize          string
}

func (s Select) SelectedValue(aliases []string) (string, error) {
	sp := selection.New(s.SelectionPrompt, aliases)

	sp.FilterPlaceholder = " " + s.FilterPlaceholder
	sp.FilterPrompt = s.FilterPrompt
	sp.Template = s.Template

	sp.ResultTemplate = ResultTemplateSelect

	pageSize, err := strconv.Atoi(s.PageSize)

	if err != nil {
		output.GetOutError("err page size must be an integer")
		return "", err
	}

	sp.PageSize = pageSize

	sp.Filter = func(filter string, choice *selection.Choice[string]) bool {
		return strings.HasPrefix(choice.Value, filter)
	}

	color := termenv.String().Foreground(termenv.ANSI256Color(206))

	sp.SelectedChoiceStyle = func(c *selection.Choice[string]) string {
		return color.Bold().Styled(c.Value)
	}

	sp.ExtendedTemplateFuncs = map[string]interface{}{
		ResultChoiceName: func(c *selection.Choice[string]) string { return c.Value },
	}

	selectedValue, err := sp.RunPrompt()

	if err != nil {
		output.GetOutError("err running prompt")
		return "", err
	}

	return selectedValue, nil
}
