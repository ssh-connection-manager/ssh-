package output

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func exitApp() {
	fmt.Println("error output")
	os.Exit(1)
}

func getOut(color *color.Color, message string, isExitApp bool) {
	_, err := color.Println(message)
	if err != nil {
		exitApp()
	}

	if isExitApp {
		os.Exit(1)
	}
}

func GetOutSuccess(successName string) {
	green := color.New(color.FgGreen)

	getOut(green, successName, !ExitApp)
}

func GetOutError(errName string) {
	red := color.New(color.FgRed)

	getOut(red, errName, ExitApp)
}
