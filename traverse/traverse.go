package traverse

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
}

func skipDirCheck(path string, info os.FileInfo) bool {
	// If the -ng flag is set, skip .git directories
	if noGit && strings.Contains(path, ".git") {
		return true
	}

	// Check if the directory is in the skip list
	if skip != "" {
		for _, skipDir := range strings.Split(skip, ",") {
			if info.IsDir() && info.Name() == skipDir {
				return true
			}
		}
	}

	return false
}

func printDirectoryTree(path string) {
	indent := strings.Repeat("  ", strings.Count(path, string(filepath.Separator))-strings.Count(dir, string(filepath.Separator)))
	if path != dir {
		fmt.Printf("%s|-- %s\n", indent, filepath.Base(path))
	} else {
		fmt.Println(path)
	}
}

func Traverse() {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if skipDirCheck(path, info) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		printDirectoryTree(path)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory %v\n", err)
	}
}
