package lib

import (
	"github.com/alecthomas/participle"
)

// The root level object that represents a go.mod file
type GoModFile struct {
	Module     string      `"module" @String`
	Statements []Statement `@@*`

	GoVersion    *string
	Requirements []Dependency
	Replacements []Replacement
	Excludes     []Exclude
}

type Statement struct {
	GoVersion    *string       `( "go" @String )`
	Requirements []Dependency  `| (("require" "(" @@* ")") | ("require" @@))`
	Replacements []Replacement `| (("replace" "(" @@* ")") | ("replace" @@))`
	Excludes     []Exclude     `| (("exclude" "(" @@* ")") | ("exclude" @@))`
}

// A struct that represents a go.mod dependency
type Dependency struct {
	ModuleName string  `@String`
	Version    string  `@Version`
	Comment    *string `("//" @String)?`
}

// A struct that represents a replace directive
type Replacement struct {
	FromModule string     `@String "=>"`
	ToModule   Dependency `@@`
}

type Exclude struct {
	Dependency Dependency `@@`
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
	ast.Flatten()
	return ast, err
}
