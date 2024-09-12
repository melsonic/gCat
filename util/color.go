package util

import (
	"fmt"
)

/// resetColor function to reset stdout text color to white
func resetColor() {
	fmt.Printf("\x1b[1;%dm", 0)
}

/// prepend ansi escape sequence for green color  
/// and return the formatted string
func getGreenText(input string) string {
	return fmt.Sprintf("\x1b[1;%dm%s", GreenColorCode, input)
}

/// print the input string in green color to stdout
/// also reset the stdout text color back to white
func PrintGreenln(input string) {
	fmt.Println(getGreenText(input))
	resetColor()
}

/// prepend ansi escape sequence for blue color  
/// and return the formatted string
func getBlueText(input string) string {
	return fmt.Sprintf("\x1b[1;%dm%s", BlueColorCode, input)
}

/// print the input string in blue color to stdout
/// also reset the stdout text color back to white
func PrintBlueln(input string) {
	fmt.Println(getBlueText(input))
	resetColor()
}

/// print an input string line with line number prepended to it 
/// in case '-n' or '-b' flag is passed
/// also increment the GLOBAL_COUNTER for next line print
func PrintBluelnWithLineNumber(lineNumber int, input string) {
	PrintBlueln(fmt.Sprintf("    %2d  %s", lineNumber, input))
	GLOBAL_COUNTER = GLOBAL_COUNTER + 1
}
