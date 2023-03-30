package main

import (
	"os"

	. "github.com/cqroot/go-snippets"
)

func main() {
	path := "./testdir/testdir"

	err := os.MkdirAll(path, os.ModePerm)
	CheckErr(err)
}
