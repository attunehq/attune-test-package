package main

import (
	"flag"
	"fmt"
	"runtime"
)

var (
	version = "1.0.1"
)

func main() {
	var showVersion = flag.Bool("version", false, "Show version information")
	var showInfo = flag.Bool("info", false, "Show system information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("attune-test-package %s\n", version)
		return
	}

	if *showInfo {
		fmt.Printf("Architecture: %s\n", runtime.GOARCH)
		fmt.Printf("OS: %s\n", runtime.GOOS)
		fmt.Printf("Go version: %s\n", runtime.Version())
		return
	}

	fmt.Printf("Hello from attune-test-package v%s!\n", version)
	fmt.Printf("Running on %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println("This is a test package for Attune CLI smoke tests.")
}