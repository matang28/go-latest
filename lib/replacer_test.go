package lib

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestReplaceVersion_Match(t *testing.T) {

	file, err := Parse(`
module github.com/matang28/go-latest
go 1.12
require (
	github.com/bla1/bla1 v1.23.1 // indirect
	github.com/bla2/bla2 v2.25.8-20190701-fuasdjhasd8
)
replace github.com/bla1/bla1 => github.com/bla1/bla1 v0.0.0-20190910155135-963c25ece259 // indirect
replace github.com/bla2/bla2 => github.com/bla2/bla2 v0.0.0-20190910155135-963c25ece259
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	expr := regexp.MustCompile("github.com/bla1")

	match := ReplaceVersion(file, expr, "latest")
	assert.True(t, match)
	assert.EqualValues(t, "latest", file.Requirements[0].Version)
	assert.EqualValues(t, "latest", file.Replacements[0].ToModule.Version)
}

func TestReplaceVersion_NoMatch(t *testing.T) {

	file, err := Parse(`
module github.com/matang28/go-latest
go 1.12
require (
	github.com/bla1/bla1 v1.23.1 // indirect
	github.com/bla2/bla2 v2.25.8-20190701-fuasdjhasd8
)
replace github.com/bla1/bla1 => github.com/bla1/bla1 v0.0.0-20190910155135-963c25ece259 // indirect
replace github.com/bla2/bla2 => github.com/bla2/bla2 v0.0.0-20190910155135-963c25ece259
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	expr := regexp.MustCompile("github.com/matang28")

	match := ReplaceVersion(file, expr, "latest")
	assert.False(t, match)
}
