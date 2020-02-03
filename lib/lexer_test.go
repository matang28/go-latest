package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexerInit(t *testing.T) {
	expected := []string{"EOF", "Parentheses", "Arrow", "String", "Version"}
	m := iniLexer.Symbols()

	assert.NotNil(t, m)
	assert.EqualValues(t, 5, len(m))

	for k := range m {
		assert.Contains(t, expected, k)
	}
}
