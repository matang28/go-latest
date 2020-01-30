package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const Latest = "latest"

func GoLatest(rootPath string, pattern *regexp.Regexp) {
	paths, err := DiscoverModFiles(rootPath)
	if err != nil {
		PrintErrorPanic(fmt.Sprintf("Failed to discover mod files on path: %s with error: %s", rootPath, err.Error()))
	}

	files := LoadFiles(paths)
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

		if ReplaceVersion(modFile, pattern, Latest) {
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
