package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("testfile")
	if err != nil {
		return
	}
	defer file.Close()

	// Write string
	nBytes, err := file.WriteString("This is a test file\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", nBytes)

	// Write []byte
	nBytes, err = file.Write([]byte("This is a test file\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", nBytes)
}
