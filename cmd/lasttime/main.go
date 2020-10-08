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
	cfg.AddTarget("Hi", 5)
	cfg.AddTarget("Link", 3)

	b, _ := json.MarshalIndent(&cfg, "", "  ")
	fmt.Println(string(b))
}
