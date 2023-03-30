package main

import (
	"fmt"
	"os"

	. "github.com/cqroot/go-snippets"
)

func WriteFile1() {
	err := os.WriteFile("testfile1", []byte("This is a test file\n"), 0o666)
	CheckErr(err)
}

func WriteFile2() {
	file, err := os.Create("testfile2")
	CheckErr(err)
	defer file.Close()

	// Write string
	nBytes, err := file.WriteString("This is a test file\n")
	CheckErr(err)
	// Output: wrote 20 bytes
	fmt.Printf("wrote %d bytes\n", nBytes)

	// Write []byte
	nBytes, err = file.Write([]byte("This is a test file\n"))
	CheckErr(err)
	// Output: wrote 20 bytes
	fmt.Printf("wrote %d bytes\n", nBytes)
}

func main() {
	WriteFile1()
	WriteFile2()
}
