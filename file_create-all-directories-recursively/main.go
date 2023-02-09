package main

import (
	"os"
)

func main() {
	path := "./testdir/testdir"

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
