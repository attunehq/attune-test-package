package main

import (
	"flag"
	"fmt"
	"runtime"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var showVersion = flag.Bool("version", false, "Show version information")
	var showInfo = flag.Bool("info", false, "Show system information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("tuna-test-package %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("built: %s\n", date)
		return
	}

	if *showInfo {
		fmt.Printf("Architecture: %s\n", runtime.GOARCH)
		fmt.Printf("OS: %s\n", runtime.GOOS)
		fmt.Printf("Go version: %s\n", runtime.Version())
		return
	}

	fmt.Printf("Hello from tuna-test-package v%s!\n", version)
	fmt.Printf("Running on %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println("This is a test package for Attune CLI smoke tests - Tuna Edition!")
}
