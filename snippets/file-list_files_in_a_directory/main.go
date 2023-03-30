package main

import (
	"fmt"
	"os"

	. "github.com/cqroot/go-snippets"
)

func main() {
	files, err := os.ReadDir("../")
	CheckErr(err)

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
