package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/cqroot/go-snippets"
)

func main() {
	file, err := os.Open("./main.go")
	CheckErr(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println("line:", fileScanner.Text())
	}
}
