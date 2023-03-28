package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir("../")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
