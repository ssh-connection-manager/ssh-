package view

import (
	"github.com/spf13/viper"
	"slices"

	"ssh+/app/output"

	"github.com/erikgeiser/promptkit/textinput"
)

type TextInput struct {
	Arguments   [][]*string
	HiddenArgs  []*string
	Placeholder string
}

func (t TextInput) currentPlaceholder(name string) string {
	return viper.GetString("Space") + name + t.Placeholder
}

func (t TextInput) DrawInput() {
	var input *textinput.TextInput
	var err error

	for _, arg := range t.Arguments {
		input = textinput.New(*arg[0])

		input.ResultTemplate = ResultTemplateTextInput

		hiddenInput := slices.Contains(t.HiddenArgs, arg[1])

		if hiddenInput {
			input.Hidden = true
		}

		input.Placeholder = t.currentPlaceholder(*arg[0])

		*arg[1], err = input.RunPrompt()

		if err != nil {
			output.GetOutError("Input error")
		}
	}
}
