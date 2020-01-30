package lib

import (
	"github.com/alecthomas/participle/lexer"
)

var iniLexer = lexer.Must(lexer.Regexp(
	`(\s+)` +
		`|(?P<Parentheses>[\(\)])` +
		`|(?P<Arrow>(=>))` +
		`|(?P<String>[a-zA-Z0-9_\-\.\/\\]*)`,
))
