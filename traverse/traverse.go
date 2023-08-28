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
	skip string
)

func init() {
	// Initialize the command-line flags
	flag.StringVar(&dir, "dir", ".", "Specify the directory to walk through. Defaults to the current directory.")
	flag.BoolVar(&noGit, "ng", false, "Skip .git directories.")
	flag.StringVar(&skip, "skip", "", "Skip comma seperated directories")
	flag.Parse()
}

func Traverse() {
	// Convert comma-separated skip directories into a slice for easy checking
	var skipDirs []string
	if skip != "" {
		skipDirs = strings.Split(skip, ",")
	}

	// Walk the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		// If the -ng flag is set, skip .git directories
		if noGit && strings.Contains(path, ".git") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip the specified directories using the -skip flag
		for _, skipDir := range skipDirs {
			if info.IsDir() && info.Name() == skipDir {
				return filepath.SkipDir
			}
		}

		// Calculate the indentation for the tree structure
		indent := strings.Repeat("  ", strings.Count(path, string(filepath.Separator))-strings.Count(dir, string(filepath.Separator)))
		if path != dir {
			fmt.Printf("%s|-- %s\n", indent, filepath.Base(path))
		} else {
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory %v\n", err)
	}
}