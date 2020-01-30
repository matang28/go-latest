package lib

import (
	"github.com/alecthomas/participle"
)

// The root level object that represents a go.mod file
type GoModFile struct {
	Module       string        `"module" @String`
	GoVersion    *string       `( "go" @String )?`
	Requirements []Dependency  `(("require" "(" @@* ")") | ("require" @@))?`
	Replacements []Replacement `@@*`
}

// A struct that represent a go.mod dependency
type Dependency struct {
	ModuleName string  `@String`
	Version    string  `@String`
	Comment    *string `("//" @String)?`
}

// A struct that represent a replace directive
type Replacement struct {
	FromModule string     `"replace" @String "=>"`
	ToModule   Dependency `@@`
}

// Will parse the given string into GoModFile struct
func Parse(source string) (*GoModFile, error) {
	p, err := participle.Build(&GoModFile{},
		participle.Lexer(iniLexer),
	)
	if err != nil {
		return nil, err
	}

	ast := &GoModFile{}
	err = p.ParseString(source, ast)
	return ast, err
}
