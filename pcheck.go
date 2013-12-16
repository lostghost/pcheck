package main

import (
	"fmt"
	"go/build"
	"os"
)

var arg string

func init() {
	// Get the package(s) from the command line
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
}

func main() {
	bp, err := build.Default.ImportDir(arg, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Printf("%+v\n", bp)
	for _, imp := range bp.Imports {
		fmt.Println(imp)
	}
	// fmt.Println("All Done")
}
