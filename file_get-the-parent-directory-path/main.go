package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(filepath.Dir(cwd))
}
