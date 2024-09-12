package main

import (
	"os"

	"github.com/melsonic/gcat/util"
)

func main() {
	/// Trim the program name present in os.Args[0]
	var cliArgs []string = os.Args[1:]
	var files []string
	var context util.Context
	for _, cliArg := range cliArgs {
		if cliArg == util.NumberFlag {
			context.SetNumberFlagState(true)
		} else if cliArg == util.NumberNonBlankFlag {
			context.SetNumberBlanklineFlagState(true)
		} else {
			files = append(files, cliArg)
		}
	}
	/// main functional component of gcat
	util.CatMainFunc(files, context)
}
