package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ModFileContent struct {
	Content []byte
	Path    string
	Error   error
}

func LoadFiles(paths []string) []ModFileContent {
	var out []ModFileContent
	for _, path := range paths {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			out = append(out, ModFileContent{Error: err, Path: path})
		} else {
			out = append(out, ModFileContent{Content: content, Path: path, Error: nil})
		}
	}

	return out
}

func DiscoverFilesByExt(rootPath string, fileSuffix string) ([]string, error) {
	var out []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return fmt.Errorf("path: %s not found or not accessible", path)
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, fileSuffix) {
			out = append(out, path)
		}
		return nil
	})
	return out, err
}

func DiscoverModFiles(rootPath string) ([]string, error) {
	return DiscoverFilesByExt(rootPath, "go.mod")
}
