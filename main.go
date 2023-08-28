package main

import (
	"app/traverse"
	"flag"
	"fmt"
)

var (
	dir   string
	noGit bool
	skip  string
)

func init() {
	// Initialize the command-line flags
	flag.StringVar(&dir, "dir", ".", "Specify the directory to walk through. Defaults to the current directory.")
	flag.BoolVar(&noGit, "ng", false, "Skip .git directories.")
	flag.StringVar(&skip, "skip", "", "Skip comma-separated directories")
	flag.Parse()

	// Set the flags in the traverse package
	traverse.SetDir(dir)
	traverse.SetNoGit(noGit)
	traverse.SetSkip(skip)
}

func main() {
	if err := traverse.Traverse(); err != nil {
		fmt.Println("Error:", err)
	}
}
