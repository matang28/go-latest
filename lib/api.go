package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const latest = "latest"

// This function will run go-latest logic against the provided root path
// and compiled regex pattern
func GoLatest(rootPath string, pattern *regexp.Regexp) {
	paths, err := discoverModFiles(rootPath)
	if err != nil {
		PrintErrorPanic(fmt.Sprintf("Failed to discover mod files on path: %s with error: %s", rootPath, err.Error()))
	}

	files := loadFiles(paths)
	for _, file := range files {
		PrintHeader(fmt.Sprintf(`File: "%s":`, file.Path))
		if file.Error != nil {
			PrintError(fmt.Sprintf(`  * Failed to read "%s" contents with error: %s`, file.Path, file.Error.Error()))
			fmt.Println()
			continue
		}

		modFile, err := Parse(string(file.Content))
		if err != nil {
			PrintError(fmt.Sprintf(`  * Failed to parse "%s" with error: %s`, file.Path, err.Error()))
			fmt.Println()
			continue
		}

		if replaceVersion(modFile, pattern, latest) {
			PrintInfo(fmt.Sprintf(`  * Will try to write the patched go.mod file to "%s".`, file.Path))
			newContent := []byte(PrintGoModFile(*modFile))
			err := ioutil.WriteFile(file.Path, newContent, os.ModePerm)
			if err != nil {
				PrintError(fmt.Sprintf(`  * Failed to write file: "%s" with the following content:\n%s`, file.Path, newContent))
				fmt.Println()
				continue
			}
			PrintInfo(fmt.Sprintf("  * File patched successfully"))
		}

		fmt.Println()
	}
}
