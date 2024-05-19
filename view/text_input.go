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

func (t TextInput) DrawInput(arguments map[string]*string) {
	var input *textinput.TextInput
	var err error

	for name, arg := range arguments {
		input = textinput.New(name)

		hiddenInput := slices.Contains(t.HiddenArgs, arg)

		if hiddenInput {
			input.Hidden = true
		}

		input.Placeholder = t.currentPlaceholder(name)

		*arg, err = input.RunPrompt()

		if err != nil {
			output.GetOutError("Ошибка ввода")
		}
	}
}
