package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type modFileContent struct {
	Content []byte
	Path    string
	Error   error
}

func loadFiles(paths []string) []modFileContent {
	var out []modFileContent
	for _, path := range paths {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			out = append(out, modFileContent{Error: err, Path: path})
		} else {
			out = append(out, modFileContent{Content: content, Path: path, Error: nil})
		}
	}

	return out
}

func discoverFilesByExt(rootPath string, fileSuffix string) ([]string, error) {
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

func discoverModFiles(rootPath string) ([]string, error) {
	return discoverFilesByExt(rootPath, "go.mod")
}
