package lib

import (
	"github.com/alecthomas/participle"
)

type GoModFile struct {
	Module       string        `"module" @String`
	GoVersion    *string       `( "go" @String )?`
	Requirements []Dependency  `(("require" "(" @@* ")") | ("require" @@))?`
	Replacements []Replacement `@@*`
}

type Dependency struct {
	ModuleName string  `@String`
	Version    string  `@String`
	Comment    *string `("//" @String)?`
}

type Replacement struct {
	FromModule string     `"replace" @String "=>"`
	ToModule   Dependency `@@`
}

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
