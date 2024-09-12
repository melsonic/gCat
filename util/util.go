package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// / entry function to the application
// / it reads the file content of each individual file from files
// / if file is a dash(-), then it reads from stdin until a EOF error is encountered
func CatMainFunc(files []string, ctx Context) {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	for _, file := range files {
		if file == DashReadStdin {
			for true {
				stdinString, scanStdinErr := readSTDIN(reader)
				if scanStdinErr == io.EOF {
					break
				}
				printLineUtil(stdinString, ctx)
			}
		} else {
			// TODO: read file and split on '\n' then PrintBluelnWithLineNumber(GLOBAL_COUNTER, line)
			var fileContent string = readFile(file)
			var fileContentLines []string = strings.Split(fileContent, NewLineCharacter)
			if fileContentLines[len(fileContentLines)-1] == EmptyString {
				fileContentLines = fileContentLines[:len(fileContentLines)-1]
			}
			for _, fileContentLine := range fileContentLines {
				printLineUtil(fileContentLine, ctx)
			}
		}
	}
	/// in case input passed through a pipe
	/// dash will be not there in files to identify
	/// read input from stdin
	stdinStat, _ := os.Stdin.Stat()
	if (stdinStat.Mode() & os.ModeCharDevice) == 0 {
		for true {
			stdinString, scanStdinErr := readSTDIN(reader)
			if scanStdinErr == io.EOF {
				break
			}
			printLineUtil(stdinString, ctx)
		}
	}
}

// / readSTDIN reads from the standard input
// / a reader variable is passed to read from standard input
// / it returns the input without the '\n' character or in
// / case of an error, it returns the error and an empty string
func readSTDIN(reader *bufio.Reader) (string, error) {
	input, scanErr := reader.ReadString('\n')
	if scanErr != nil {
		return EmptyString, scanErr
	}
	input = input[:len(input)-1]
	return input, nil
}

// / readFile reads the content of a file
// / in case of invalid path, it throws an file not found error
func readFile(fileName string) string {
	fileData, fileReadError := os.ReadFile(fileName)
	if fileReadError != nil {
		log.Fatalf("file -> %s not found\n", fileName)
	}
	return string(fileData)
}

// / printLineUtil prints a text line based on the set context
// / if '-n' flag is set it attaches a line number in front of the line
// / in case of '-b' flag also, it prepends line number except for empty string
func printLineUtil(line string, ctx Context) {
	if !ctx.GetNumberBlanklineFlagState() && !ctx.GetNumberFlagState() {
		PrintBlueln(line)
		return
	}
	if line == EmptyString && ctx.GetNumberBlanklineFlagState() {
		PrintBlueln(line)
		return
	}
	PrintBluelnWithLineNumber(GLOBAL_COUNTER, line)
}
