package main

import (
	"fmt"
	"os"
)

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func main() {
	fmt.Println(fileExists("./main.go"))
	fmt.Println(fileExists("./test.go"))
}
