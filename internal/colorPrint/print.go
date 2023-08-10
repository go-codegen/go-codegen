package colorPrint

import (
	"fmt"
	"os"
)

const (
	redColor    = "\033[31m"
	greenColor  = "\033[32m"
	yellowColor = "\033[33m"
	resetColor  = "\033[0m"
)

func printColoredError(message string) {
	fmt.Fprintf(os.Stderr, "%s%s%s\n", redColor, message, resetColor)
}

func printColoredSuccess(message string) {
	fmt.Fprintf(os.Stdout, "%s%s%s\n", greenColor, message, resetColor)
}

func printColoredWarning(message string) {
	fmt.Fprintf(os.Stdout, "%s%s%s\n", yellowColor, message, resetColor)
}

func PrintError(err error) {
	printColoredError(err.Error())
}

func PrintSuccess(message string) {
	printColoredSuccess(message)
}

func PrintWarning(message string) {
	printColoredWarning(message)
}
