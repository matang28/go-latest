package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiscoverFilesByExt_FoundModFiles(t *testing.T) {
	files, err := DiscoverFilesByExt("../test_data/", "go.mod")
	assert.Nil(t, err)
	assert.NotNil(t, files)

	assert.EqualValues(t, 4, len(files))
	assert.Contains(t, files, "../test_data/go.mod")
	assert.Contains(t, files, "../test_data/one/go.mod")
	assert.Contains(t, files, "../test_data/one/one_sub/go.mod")
	assert.Contains(t, files, "../test_data/two/go.mod")
}

func TestDiscoverFilesByExt_NotFoundModFiles(t *testing.T) {
	files, err := DiscoverFilesByExt("./", "go.mod")
	assert.Nil(t, err)
	assert.Empty(t, files)
}

func TestDiscoverFilesByExt_InvalidPath(t *testing.T) {
	files, err := DiscoverFilesByExt("/madeup/path/to/nothing", "go.mod")
	assert.NotNil(t, err)
	assert.Empty(t, files)
}
