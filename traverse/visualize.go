package traverse

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	yellow = "\033[33m"
	green  = "\033[32m"
	reset  = "\033[0m"
)

func printDirectoryTree(path string, info os.FileInfo) {
	indent := strings.Repeat("  ", strings.Count(path, string(filepath.Separator))-strings.Count(dir, string(filepath.Separator)))
	if path != dir {
		if info.IsDir() {
			// If it's a directory, print in yellow
			fmt.Printf("%s|-- %s%s%s\n", indent, yellow, filepath.Base(path), reset)
		} else {
			// If it's a file, print in green
			fmt.Printf("%s|-- %s%s%s\n", indent, green, filepath.Base(path), reset)
		}
	} else {
		fmt.Println(path)
	}
}
