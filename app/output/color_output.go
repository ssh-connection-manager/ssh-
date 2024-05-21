package output

import (
	"os"

	"github.com/fatih/color"
)

func GetOutSuccess(successName string) {
	green := color.New(color.BgGreen)
	green.Println(successName)
}

func GetOutError(errName string) {
	red := color.New(color.FgRed)
	red.Println(errName)

	os.Exit(1)
}
