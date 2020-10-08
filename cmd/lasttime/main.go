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
var pathStr string

func parseOption() bool {
	var optVersion bool

	flag.BoolVar(&optVersion, "v", false, "show version")
	flag.BoolVar(&optVersion, "version", false, "show version")
	fileptr := flag.String("f", "", "input file")
	flag.Parse()
	if optVersion {
		fmt.Printf("version:%v(%v)", version, revision)
		return false
	}
	if fileptr != nil {
		pathStr = *fileptr
	}
	return true
}

func main() {
	if !parseOption() {
		return
	}

	lasttime.ExpBufio(5, true)

	if pathStr != "" {
		fmt.Println("file is specified:", pathStr)
		reader, err := lasttime.ExpGetURL(pathStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer reader.Close()

		err = lasttime.ExpSaveFile(reader)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Writer Success")
	} else {
		fmt.Println("file is not specified")
	}

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
