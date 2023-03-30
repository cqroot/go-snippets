package main

import (
	"fmt"
	"os"

	. "github.com/cqroot/go-snippets"
)

func main() {
	content, err := os.ReadFile("./main.go")
	CheckErr(err)

	fmt.Println(string(content))
}
