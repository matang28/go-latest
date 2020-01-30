package lib

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

var errColor = color.New(color.FgRed)
var errColorItal = color.New(color.FgHiRed).Add(color.Italic).Add(color.Bold)
var infColor = color.New(color.FgWhite)
var headerColor = color.New(color.FgYellow).Add(color.Underline)

func PrintError(message string) {
	errColor.Println(message)
}

func PrintErrorCommand(message string) {
	errColorItal.Print(message)
}

func PrintErrorPanic(message string) {
	errColor.Println(message)
	os.Exit(1)
}

func PrintInfo(message string) {
	infColor.Println(message)
}

func PrintHeader(message string) {
	headerColor.Println(message)
}

func PrintGoModFile(file GoModFile) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("module %s\n", file.Module))

	if file.GoVersion != nil {
		sb.WriteString(fmt.Sprintf("go %s\n", *file.GoVersion))
	}

	if file.Requirements != nil {
		sb.WriteString("require (\n")
	}

	for _, req := range file.Requirements {
		sb.WriteString(fmt.Sprintf("\t%s %s", req.ModuleName, req.Version))
		if req.Comment != nil {
			sb.WriteString(fmt.Sprintf(" // %s\n", *req.Comment))
		} else {
			sb.WriteString("\n")
		}
	}

	if file.Requirements != nil {
		sb.WriteString(")\n")
	}

	for _, rep := range file.Replacements {
		sb.WriteString(fmt.Sprintf("replace %s => %s %s", rep.FromModule, rep.ToModule.ModuleName, rep.ToModule.Version))
		if rep.ToModule.Comment != nil {
			sb.WriteString(fmt.Sprintf(" // %s\n", *rep.ToModule.Comment))
		} else {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
