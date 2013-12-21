package main

import (
	"flag"
	"fmt"
	"go/build"
	"os"
)

// Command line args
var args []string

// Working directory
var pwd = ""

// Output streams
var (
	stdout = os.Stdout
	stderr = os.Stderr
)

// Build context
var buildContext = build.Default

// Configuration
var config struct {
	includeTests bool
}

func init() {
	// Parse command line flags
	flag.BoolVar(&config.includeTests, "tests", false, "Include test packages?")
	flag.Parse()

	// Parse command line args
	args = flag.Args()

	// Get pwd
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		fmt.Fprintf(stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Fprintf(stdout, "Include tests? %+v\n", config.includeTests)

	bp, err := buildContext.Import(args[0], pwd, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Printf("%+v\n", bp)
	for _, imp := range bp.Imports {
		fmt.Println(imp)
	}
	if config.includeTests {
		for _, imp := range bp.TestImports {
			fmt.Println(imp)
		}
	}
}
