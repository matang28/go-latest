package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintGoModFile(t *testing.T) {
	content := `module github.com/matang28/go-latest
go 1.12
require (
	github.com/bla1/bla1 v1.23.1 // indirect
	github.com/bla2/bla2 v2.25.8-20190701-fuasdjhasd8
)
replace github.com/bla1/bla1 => github.com/bla1/bla1 v0.0.0-20190910155135-963c25ece259 // indirect
replace github.com/bla2/bla2 => github.com/bla2/bla2 v0.0.0-20190910155135-963c25ece259
`

	file, err := Parse(content)

	assert.Nil(t, err)

	printed := PrintGoModFile(*file)
	assert.NotEmpty(t, printed)
	assert.EqualValues(t, content, printed)
}
