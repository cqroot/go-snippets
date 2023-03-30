package main

import (
	"fmt"
	"os"
)

func ReadFile(name string) {
	content, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func main() {
	ReadFile("./main.go")
}
