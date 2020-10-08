package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/hidenden/lasttime"
)

var revision string

func main() {
	fmt.Println("Hello World")
	fmt.Println("Number of Go routine", runtime.NumGoroutine())
	fmt.Println("Command revision", revision)

	cfg := lasttime.Config{
		Mode: "Hoge",
	}
	//{Name: "World", Threshold:2}}

	b, _ := json.MarshalIndent(&cfg, "", "  ")
	fmt.Println(string(b))
}
