package lib

import (
	"os/exec"
)

func goModTidy(path string) (response string, err error) {
	c := exec.Command("go", "mod", "tidy")
	c.Dir = path
	bytes, err := c.CombinedOutput()
	return string(bytes), err
}
