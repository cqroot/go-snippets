package main

import (
	"errors"
	"os"
)

func main() {
	path := "./testdir"

	// Check for the directory’s existence first.
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}
