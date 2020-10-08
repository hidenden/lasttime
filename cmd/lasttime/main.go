package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var revision string

func main() {
	fmt.Println("Hello World")
	fmt.Println("Number of Go routine", runtime.NumGoroutine())
	fmt.Println("Command revision", revision)

	var cfg config

	b, _ := json.MarshalIndent(&cfg, "", "  ")
	fmt.Println(string(b))
}
