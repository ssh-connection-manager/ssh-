package view

import (
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
	"slices"

	"github.com/erikgeiser/promptkit/textinput"
)

type TextInput struct {
	Arguments   [][]*string
	HiddenArgs  []*string
	Placeholder string
}

func (t TextInput) currentPlaceholder(name string) (string, error) {
	return " " + name + t.Placeholder, nil
}

func (t TextInput) DrawInput() error {
	var input *textinput.TextInput
	var err error

	for _, arg := range t.Arguments {
		input = textinput.New(*arg[0])

		input.ResultTemplate = ResultTemplateTextInput

		hiddenInput := slices.Contains(t.HiddenArgs, arg[1])

		if hiddenInput {
			input.Hidden = true
		}

		input.Placeholder, _ = t.currentPlaceholder(*arg[0])
		*arg[1], err = input.RunPrompt()

		if err != nil {
			output.GetOutError("err run prompt this current args")
			return err
		}
	}

	return nil
}
