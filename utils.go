package dir

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	// doubleBackslash defines windows path separator
	doubleBackslash = "\\"
	// separator defines mac or linux path separator
	separator = "/"
)

// Replace mismatch separator to current os separator
func Replace(path string) string {
	if path = strings.TrimSpace(path); path == "" {
		return path
	} else if string(path[0]) == Home {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		path = filepath.Join(homeDir, path[1:])
	}
	return strings.ReplaceAll(path, doubleBackslash, separator)
}
