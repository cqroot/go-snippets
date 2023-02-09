package main

import "fmt"

type LogLevel int

const (
	// +1 for having numbers start at 1 instead of 0
	Info LogLevel = iota + 1
	Warn
	Error
)

func main() {
	// Output: 1 2 3
	fmt.Println(Info, Warn, Error)
}
