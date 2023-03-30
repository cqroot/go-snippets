package main

import (
	"os"

	. "github.com/cqroot/go-snippets"
)

func main() {
	oldLocation := "./testfile_file_to_move"
	newLocation := "./testfile"

	err := os.WriteFile(oldLocation, []byte("This is the file to move.\n"), 0o666)
	CheckErr(err)

	err = os.Rename(oldLocation, newLocation)
	CheckErr(err)
}
