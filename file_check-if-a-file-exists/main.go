package main

import (
	"fmt"
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func main() {
	fmt.Println(fileExists("./main.go"))
	fmt.Println(fileExists("./test.go"))
}
