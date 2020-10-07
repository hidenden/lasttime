package main

import (
	"fmt"
	"runtime"
)

var revision string

func main() {
	fmt.Println("Hello World")
	fmt.Println("Number of Go routine", runtime.NumGoroutine())
	fmt.Println("Command revision", revision)
}
