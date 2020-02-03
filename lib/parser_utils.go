package lib

import "reflect"

func (gmf *GoModFile) Flatten() {
	for _, stmt := range gmf.Statements {
		if isIt(stmt.Requirements) {
			gmf.Requirements = append(gmf.Requirements, stmt.Requirements...)
		}

		if isIt(stmt.Replacements) {
			gmf.Replacements = append(gmf.Replacements, stmt.Replacements...)
		}

		if isIt(stmt.Excludes) {
			gmf.Excludes = append(gmf.Excludes, stmt.Excludes...)
		}

		if stmt.GoVersion != nil {
			gmf.GoVersion = new(string)
			*gmf.GoVersion = *stmt.GoVersion
		}
	}
}

func isIt(statement interface{}) bool {
	if statement != nil {
		valueOf := reflect.ValueOf(statement)
		if valueOf.Kind() == reflect.Array || valueOf.Kind() == reflect.Slice {
			if valueOf.Len() > 0 {
				return true
			}
		}
	}
	return false
}
