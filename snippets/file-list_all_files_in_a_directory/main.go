package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("../", func(name string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if info.IsDir() {
			return nil
		}

		fmt.Println(name)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
