package main

import (
	"fmt"
	"os"

	. "github.com/cqroot/go-snippets"
)

func main() {
	entries, err := os.ReadDir("../")
	CheckErr(err)

	for _, entry := range entries {
		fmt.Println(entry.Name(), entry.IsDir())
	}
}
