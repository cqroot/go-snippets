package main

import (
	"fmt"
	"os"

	. "github.com/cqroot/go-snippets"
)

func main() {
	cwd, err := os.Getwd()
	CheckErr(err)

	fmt.Println(cwd)
}
