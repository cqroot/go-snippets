package main

import (
	"fmt"
	"os/exec"

	. "github.com/cqroot/go-snippets"
)

func main() {
	outb, err := exec.Command("/bin/sh", "-c", "echo 'test message' | awk '{print $2}'").Output()
	CheckErr(err)

	fmt.Printf("Output: {%s}\n", outb)
}
