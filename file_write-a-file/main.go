package main

import (
	"fmt"
	"os"
)

func WriteFile1() {
	err := os.WriteFile("testfile1", []byte("This is a test file\n"), 0666)
	if err != nil {
		panic(err)
	}
}

func WriteFile2() {
	file, err := os.Create("testfile2")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write string
	nBytes, err := file.WriteString("This is a test file\n")
	if err != nil {
		panic(err)
	}
	// output: wrote 20 bytes
	fmt.Printf("wrote %d bytes\n", nBytes)

	// Write []byte
	nBytes, err = file.Write([]byte("This is a test file\n"))
	if err != nil {
		panic(err)
	}
	// output: wrote 20 bytes
	fmt.Printf("wrote %d bytes\n", nBytes)
}

func main() {
	WriteFile1()
	WriteFile2()
}
