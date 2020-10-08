package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"runtime"

	"github.com/hidenden/lasttime"
)

var revision string
var version string

func parseOption() bool {
	var optVersion bool
	flag.BoolVar(&optVersion, "v", false, "show version")
	flag.BoolVar(&optVersion, "version", false, "show version")
	flag.Parse()
	if optVersion {
		fmt.Printf("version:%v(%v)", version, revision)
		return false
	}
	return true
}

func main() {
	if !parseOption() {
		return
	}

	lasttime.ExpBufio(5, true)

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
