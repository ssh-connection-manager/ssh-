package view

import (
	"os"
	"slices"

	"ssh+/app/output"

	"github.com/erikgeiser/promptkit/textinput"
)

type TextInput struct {
	HiddenArgs  []*string
	Placeholder string
}

func (t TextInput) currentPlaceholder(name string) string {
	return os.Getenv("SPACE") + name + t.Placeholder
}

func (t TextInput) DrawInput(arguments [][]*string) {
	var input *textinput.TextInput
	var err error

	for _, arg := range arguments {
		input = textinput.New(*arg[0])

		hiddenInput := slices.Contains(t.HiddenArgs, arg[1])

		if hiddenInput {
			input.Hidden = true
		}

		input.Placeholder = t.currentPlaceholder(*arg[0])

		*arg[1], err = input.RunPrompt()

		if err != nil {
			output.GetOutError("Ошибка ввода")
		}
	}
}
