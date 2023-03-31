package main

import (
	"errors"
	"os"

	. "github.com/cqroot/go-snippets"
)

func CreateDir(dir string) error {
	// Check for the directory’s existence first.
	_, err := os.Stat(dir)

	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		return err
	}

	return err
}

func main() {
	CheckErr(CreateDir("./testdir"))
}
