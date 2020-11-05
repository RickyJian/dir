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

// Diff returns move and override directories and files
func Diff(dest, src []string) ([]string, []string) {
	if len(src) == 0 {
		return []string{}, []string{}
	} else if len(dest) == 0 {
		return src, []string{}
	}

	destSet := make(map[string]struct{})
	for _, d := range dest {
		destSet[d] = struct{}{}
	}
	var plus, minus []string
	for _, s := range src {
		if _, ok := destSet[s]; ok {
			minus = append(minus, s)
		} else {
			plus = append(plus, s)
		}
	}
	return plus, minus
}
