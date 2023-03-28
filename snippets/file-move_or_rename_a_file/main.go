package main

import (
	"os"
)

func main() {
	oldLocation := "./testfile_file_to_move"
	newLocation := "./testfile"

	err := os.WriteFile(oldLocation, []byte("This is the file to move.\n"), 0666)
	if err != nil {
		panic(err)
	}

	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		panic(err)
	}
}
