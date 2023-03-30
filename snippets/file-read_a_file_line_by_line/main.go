package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println("line:", fileScanner.Text())
	}
}

func main() {
	ReadFile("./main.go")
}
