package traverse

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
)

var (
	dir   string
	noGit bool
	skip  string
	gitStatus bool
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



func SetGitStatus(gs bool) {
    gitStatus = gs
}

func getGitStatus(path string) (status string, err error) {
	cmd := exec.Command("git", "ls-files", "--other", "--modified", "--exclude-standard", path)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	if len(output) > 0 {
		return "Untracked or Modified", nil
	}

	cmd = exec.Command("git", "log", "-1", "--format=%h - %cr", path)
	output, err = cmd.Output()
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(string(output)), nil
}