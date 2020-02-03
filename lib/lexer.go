package lib

import (
	"github.com/alecthomas/participle/lexer"
)

var iniLexer = lexer.Must(lexer.Regexp(
	`([\s\n\r\t]+)` +
		`|(?P<Parentheses>[\(\)])` +
		`|(?P<Arrow>(=>))` +
		`|(?P<Version>[v][a-zA-Z0-9_\+\.@\-\/]+)` +
		`|(?P<String>[a-zA-Z0-9_\+\.@\-\/]+)`,
))
