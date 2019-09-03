package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Load ...
func Load(file string) (string, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0600)
	if err != nil {
		return "", fmt.Errorf("Can not open file: %s, %v", file, err)
	}

	defer f.Close()

	return GetText(f)
}

// GetText from file (reader)
func GetText(r io.Reader) (string, error) {
	v, _ := ioutil.ReadAll(r)

	return string(v), nil
}
