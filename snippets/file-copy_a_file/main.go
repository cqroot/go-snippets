package main

import (
	"fmt"
	"io"
	"os"

	. "github.com/cqroot/go-snippets"
)

func CopyFile(src, dst string) (int64, error) {
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !srcFileInfo.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {
	nBytes, err := CopyFile("./main.go", "testfile")
	CheckErr(err)

	// This value is equal to the value output by `du -b main.go`.
	fmt.Println(nBytes)
}
