package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse_JustModule(t *testing.T) {
	file, err := Parse("module github.com/matang28/go-latest")
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.EqualValues(t, "github.com/matang28/go-latest", file.Module)
	assert.Nil(t, file.GoVersion)
	assert.Nil(t, file.Requirements)
	assert.Nil(t, file.Replacements)
}

func TestParse_ModuleAndVersion(t *testing.T) {
	file, err := Parse(`
module github.com/matang28/go-latest
go 1.12
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.EqualValues(t, "github.com/matang28/go-latest", file.Module)
	assert.EqualValues(t, "1.12", *file.GoVersion)
	assert.Nil(t, file.Requirements)
	assert.Nil(t, file.Replacements)
}

func TestParse_WithEmptyRequirements(t *testing.T) {
	file, err := Parse(`
module github.com/matang28/go-latest
go 1.12

require (
)
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.EqualValues(t, "github.com/matang28/go-latest", file.Module)
	assert.EqualValues(t, "1.12", *file.GoVersion)
	assert.Nil(t, file.Requirements)
	assert.Nil(t, file.Replacements)
}

func TestParse_WithSingleRequirement(t *testing.T) {
	file, err := Parse(`
module github.com/matang28/go-latest
go 1.12

require github.com/bla/bla v1.23.1 // indirect
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.EqualValues(t, "github.com/matang28/go-latest", file.Module)
	assert.EqualValues(t, "1.12", *file.GoVersion)
	assert.EqualValues(t, 1, len(file.Requirements))
	assert.EqualValues(t, "github.com/bla/bla", file.Requirements[0].ModuleName)
	assert.EqualValues(t, "v1.23.1", file.Requirements[0].Version)
	assert.EqualValues(t, "indirect", *file.Requirements[0].Comment)
	assert.Nil(t, file.Replacements)
}

func TestParse_WithMultipleRequirements(t *testing.T) {
	file, err := Parse(`
module github.com/matang28/go-latest
go 1.12

require (
	github.com/bla1/bla1 v1.23.1 // indirect
	github.com/bla2/bla2 v2.25.8-20190701-fuasdjhasd8
)
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.EqualValues(t, "github.com/matang28/go-latest", file.Module)
	assert.EqualValues(t, "1.12", *file.GoVersion)
	assert.EqualValues(t, 2, len(file.Requirements))

	assert.EqualValues(t, "github.com/bla1/bla1", file.Requirements[0].ModuleName)
	assert.EqualValues(t, "v1.23.1", file.Requirements[0].Version)
	assert.EqualValues(t, "indirect", *file.Requirements[0].Comment)

	assert.EqualValues(t, "github.com/bla2/bla2", file.Requirements[1].ModuleName)
	assert.EqualValues(t, "v2.25.8-20190701-fuasdjhasd8", file.Requirements[1].Version)
	assert.Nil(t, file.Requirements[1].Comment)

	assert.Nil(t, file.Replacements)
}

func TestParse_WithReplacements(t *testing.T) {
	file, err := Parse(`
module github.com/matang28/go-latest

replace github.com/bla1/bla1 => github.com/bla1/bla1 v0.0.0-20190910155135-963c25ece259 // indirect
require (
	github.com/bla1/bla1 v1.23.1 // indirect
	github.com/bla2/bla2 v2.25.8-20190701-fuasdjhasd8
)
replace github.com/bla2/bla2 => github.com/bla2/bla2 v0.0.0-20190910155135-963c25ece259
go 1.12
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.EqualValues(t, "github.com/matang28/go-latest", file.Module)
	assert.EqualValues(t, "1.12", *file.GoVersion)
	assert.EqualValues(t, 2, len(file.Requirements))
	assert.EqualValues(t, 2, len(file.Replacements))

	assert.EqualValues(t, "github.com/bla1/bla1", file.Requirements[0].ModuleName)
	assert.EqualValues(t, "v1.23.1", file.Requirements[0].Version)
	assert.EqualValues(t, "indirect", *file.Requirements[0].Comment)

	assert.EqualValues(t, "github.com/bla2/bla2", file.Requirements[1].ModuleName)
	assert.EqualValues(t, "v2.25.8-20190701-fuasdjhasd8", file.Requirements[1].Version)
	assert.Nil(t, file.Requirements[1].Comment)

	assert.EqualValues(t, "github.com/bla1/bla1", file.Replacements[0].FromModule)
	assert.EqualValues(t, "github.com/bla1/bla1", file.Replacements[0].ToModule.ModuleName)
	assert.EqualValues(t, "v0.0.0-20190910155135-963c25ece259", file.Replacements[0].ToModule.Version)
	assert.EqualValues(t, "indirect", *file.Replacements[0].ToModule.Comment)

	assert.EqualValues(t, "github.com/bla2/bla2", file.Replacements[1].FromModule)
	assert.EqualValues(t, "github.com/bla2/bla2", file.Replacements[1].ToModule.ModuleName)
	assert.EqualValues(t, "v0.0.0-20190910155135-963c25ece259", file.Replacements[1].ToModule.Version)
	assert.Nil(t, file.Replacements[1].ToModule.Comment)
}

func TestParse_WithMultipleReplacements(t *testing.T) {
	file, err := Parse(`
module github.com/matang28/go-latest

replace github.com/bla1/bla1 => github.com/bla1/bla1 v0.0.0-20190910155135-963c25ece259 // indirect
require (
	github.com/bla1/bla1 v1.23.1 // indirect
	github.com/bla2/bla2 v2.25.8-20190701-fuasdjhasd8
)
replace (
	github.com/bla2/bla2 => github.com/bla2/bla2 v0.0.0-20190910155135-963c25ece259
	github.com/bla3/bla3 => github.com/bla3/bla3 v0.0.0-20190910155135-963c25ece259
)
go 1.12
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.EqualValues(t, "github.com/matang28/go-latest", file.Module)
	assert.EqualValues(t, "1.12", *file.GoVersion)
	assert.EqualValues(t, 2, len(file.Requirements))
	assert.EqualValues(t, 3, len(file.Replacements))

	assert.EqualValues(t, "github.com/bla1/bla1", file.Requirements[0].ModuleName)
	assert.EqualValues(t, "v1.23.1", file.Requirements[0].Version)
	assert.EqualValues(t, "indirect", *file.Requirements[0].Comment)

	assert.EqualValues(t, "github.com/bla2/bla2", file.Requirements[1].ModuleName)
	assert.EqualValues(t, "v2.25.8-20190701-fuasdjhasd8", file.Requirements[1].Version)
	assert.Nil(t, file.Requirements[1].Comment)

	assert.EqualValues(t, "github.com/bla1/bla1", file.Replacements[0].FromModule)
	assert.EqualValues(t, "github.com/bla1/bla1", file.Replacements[0].ToModule.ModuleName)
	assert.EqualValues(t, "v0.0.0-20190910155135-963c25ece259", file.Replacements[0].ToModule.Version)
	assert.EqualValues(t, "indirect", *file.Replacements[0].ToModule.Comment)

	assert.EqualValues(t, "github.com/bla2/bla2", file.Replacements[1].FromModule)
	assert.EqualValues(t, "github.com/bla2/bla2", file.Replacements[1].ToModule.ModuleName)
	assert.EqualValues(t, "v0.0.0-20190910155135-963c25ece259", file.Replacements[1].ToModule.Version)
	assert.Nil(t, file.Replacements[1].ToModule.Comment)

	assert.EqualValues(t, "github.com/bla3/bla3", file.Replacements[2].FromModule)
	assert.EqualValues(t, "github.com/bla3/bla3", file.Replacements[2].ToModule.ModuleName)
	assert.EqualValues(t, "v0.0.0-20190910155135-963c25ece259", file.Replacements[2].ToModule.Version)
	assert.Nil(t, file.Replacements[2].ToModule.Comment)
}

func TestParse_WithUncommonValues(t *testing.T) {
	file, err := Parse(`
module github.com/matang28/go-latest

require (
	github.com/bla1/bla1 v1.23.1+incompatible
	github.com/bla2/bla2 v2.25.8-20190701-fuasdjhasd8
	github.com/bla2/bla3 latest
)
replace (
	github.com/bla2/bla2 v1.2.3 => github.com/bla2/bla2 v0.0.0-20190910155135-963c25ece259
)
go 1.12
`)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.EqualValues(t, "github.com/matang28/go-latest", file.Module)
	assert.EqualValues(t, "1.12", *file.GoVersion)
	assert.EqualValues(t, 3, len(file.Requirements))
	assert.EqualValues(t, 1, len(file.Replacements))

	assert.EqualValues(t, "github.com/bla1/bla1", file.Requirements[0].ModuleName)
	assert.EqualValues(t, "v1.23.1+incompatible", file.Requirements[0].Version)
	assert.Nil(t, file.Requirements[0].Comment)

	assert.EqualValues(t, "github.com/bla2/bla2", file.Requirements[1].ModuleName)
	assert.EqualValues(t, "v2.25.8-20190701-fuasdjhasd8", file.Requirements[1].Version)
	assert.Nil(t, file.Requirements[1].Comment)

	assert.EqualValues(t, "github.com/bla2/bla3", file.Requirements[2].ModuleName)
	assert.EqualValues(t, "latest", file.Requirements[2].Version)
	assert.Nil(t, file.Requirements[2].Comment)

	assert.EqualValues(t, "github.com/bla2/bla2", file.Replacements[0].FromModule)
	assert.EqualValues(t, "github.com/bla2/bla2", file.Replacements[0].ToModule.ModuleName)
	assert.EqualValues(t, "v0.0.0-20190910155135-963c25ece259", file.Replacements[0].ToModule.Version)
}
