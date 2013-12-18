package main

import (
	"flag"
	"fmt"
	"go/build"
	"os"
)

var args []string
var pwd = ""
var stdout = os.Stdout
var stderr = os.Stderr
var buildContext = build.Default

func init() {
	// Get the package(s) from the command line
	flag.Parse()
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
	bp, err := buildContext.Import(args[0], pwd, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Printf("%+v\n", bp)
	for _, imp := range bp.Imports {
		fmt.Println(imp)
	}
}
