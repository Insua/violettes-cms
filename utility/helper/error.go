package helper

import (
	"github.com/fatih/color"
	"os"
)

func ConsoleErrorMessage(msg string, err error) {
	color.Red(msg)
	if err != nil {
		color.Red("error message: %s", err.Error())
	}
	os.Exit(1)
}

func ConsoleErrorMessageExit(msg string, err error) {
	ConsoleErrorMessage(msg, err)
	os.Exit(1)
}
