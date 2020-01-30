package lib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestDiscoverFilesByExt_FoundModFiles(t *testing.T) {
	files, err := discoverFilesByExt(filepath.Join("..", "test_data"), "go.mod")
	assert.Nil(t, err)
	assert.NotNil(t, files)

	assert.EqualValues(t, 4, len(files))
	assert.Contains(t, files, filepath.Join("..", "test_data", "go.mod"))
	assert.Contains(t, files, filepath.Join("..", "test_data", "one", "go.mod"))
	assert.Contains(t, files, filepath.Join("..", "test_data", "one", "one_sub", "go.mod"))
	assert.Contains(t, files, filepath.Join("..", "test_data", "two", "go.mod"))
}

func TestDiscoverFilesByExt_NotFoundModFiles(t *testing.T) {
	pwd, err := os.Getwd()
	assert.Nil(t, err)

	files, err := discoverFilesByExt(pwd, "go.mod")
	assert.Nil(t, err)
	assert.Empty(t, files)
}

func TestDiscoverFilesByExt_InvalidPath(t *testing.T) {
	files, err := discoverFilesByExt("/madeup/path/to/nothing", "go.mod")
	assert.NotNil(t, err)
	assert.Empty(t, files)
}
