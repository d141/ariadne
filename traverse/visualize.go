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
		statusStr := ""

		if gitStatus && !info.IsDir() {
			status, err := getGitStatus(path)
			if err == nil && status != "" {
				statusStr = fmt.Sprintf(" [%s]", status)
			}
		}
		
		if info.IsDir() {
			// If it's a directory, print in blue
			fmt.Printf("%s|-- %s%s%s%s\n", indent, yellow, filepath.Base(path), reset, statusStr)
		} else {
			// If it's a file, print in green
			fmt.Printf("%s|-- %s%s%s%s\n", indent, green, filepath.Base(path), reset, statusStr)
		}
	} else {
		fmt.Println(path)
	}
}