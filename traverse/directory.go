package traverse

import (
	"os"
	"strings"
)

func skipDirCheck(path string, info os.FileInfo) bool {
	// If the noGit flag is set, skip .git directories
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
