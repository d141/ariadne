package traverse

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	dir   string
	noGit bool
	skip  string
)

// SetDir sets the directory
func SetDir(d string) {
	dir = d
}

// SetNoGit sets the noGit flag
func SetNoGit(ng bool) {
	noGit = ng
}

// SetSkip sets the skip flag
func SetSkip(s string) {
	skip = s
}

func Traverse() error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Error accessing path %q: %v", path, err)
		}

		if skipDirCheck(path, info) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		printDirectoryTree(path, info)
		return nil
	})
}
