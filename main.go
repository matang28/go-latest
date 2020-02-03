package main

import (
	"flag"
	"github.com/matang28/go-latest/lib"
	"os"
	"regexp"
)

var tidyFlag = flag.Bool("tidy", false, "adding this flag will run go mod tidy on detected files")

func main() {
	args := os.Args[1:3]
	os.Args = os.Args[2:]
	flag.Parse()

	if len(args) != 2 {
		lib.PrintError(`GoLatest expects exactly 2 arguments: "go-latest "<MATCH EXPRESSION>" <ROOT FOLDER>". For example:`)
		lib.PrintErrorCommand(`go-latest ".*" .   `)
		lib.PrintErrorPanic("will go over all go.mod files in current directory and update any dependency to the latest version.")
	}

	patternStr := args[0]
	rootFolder := args[1]

	var err error
	if rootFolder == "." {
		rootFolder, err = os.Getwd()
		if err != nil {
			lib.PrintErrorPanic(err.Error())
		}
	}

	pattern := regexp.MustCompile(patternStr)

	files := lib.GoLatest(rootFolder, pattern)

	if *tidyFlag {
		lib.GoTidy(files)
	}
}
