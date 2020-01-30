package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexerInit(t *testing.T) {
	expected := []string{"EOF", "Parentheses", "Arrow", "String"}
	m := iniLexer.Symbols()

	assert.NotNil(t, m)
	assert.EqualValues(t, 4, len(m))

	for k := range m {
		assert.Contains(t, expected, k)
	}
}
