package main

import (
	"fmt"
	"os/exec"
)

func main() {
	outb, err := exec.Command("/bin/sh", "-c", "echo 'test message' | awk '{print $2}'").Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Output: {%s}\n", outb)
}
