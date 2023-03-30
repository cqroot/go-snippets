package main

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/cqroot/go-snippets"
)

func main() {
	err := filepath.Walk("../", func(name string, info os.FileInfo, err error) error {
		CheckErr(err)

		if info.IsDir() {
			return nil
		}

		fmt.Println(name)
		return nil
	})
	CheckErr(err)
}
