package main

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/cqroot/go-snippets"
)

func main() {
	cwd, err := os.Getwd()
	CheckErr(err)

	fmt.Println(filepath.Dir(cwd))
}
