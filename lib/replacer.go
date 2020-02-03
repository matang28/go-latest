package lib

import (
	"fmt"
	"regexp"
)

func replaceVersion(file *GoModFile, pattern *regexp.Regexp, with string) bool {
	var anyMatched = false

	for idx := range file.Requirements {
		if pattern.MatchString(file.Requirements[idx].ModuleName) {
			PrintInfo(fmt.Sprintf(`  * Replacing "%s" from version "%s" to version "%s" (require)`,
				file.Requirements[idx].ModuleName, file.Requirements[idx].Version, with))
			file.Requirements[idx].Version = with
			anyMatched = true
		}
	}

	for idx := range file.Replacements {
		if pattern.MatchString(file.Replacements[idx].ToModule.ModuleName) {
			file.Replacements[idx].ToModule.Version = with
			anyMatched = true
			PrintInfo(fmt.Sprintf(`  * Replacing "%s" from version "%s" to version "%s" (replace)`,
				file.Replacements[idx].ToModule.ModuleName, file.Replacements[idx].ToModule.Version, with))
		}
	}

	return anyMatched
}
