package dir

import (
	"errors"
	"os"
	"strings"
)

const (
	// windowsSeparator defines windows path separator
	windowsSeparator = "\\"
	// separator defines mac or linux path separator
	separator = "/"
)

var (
	// PathSeparator defines string type path separator
	PathSeparator = string(os.PathSeparator)
)

// Dir defines basic directory struct
type Dir struct {
	Path       string
	Nodes      []string
	Filename   string
	Files      []string
	Permission string
}

// New new dir
func New(path string) *Dir {
	path = replace(path)
	if pathLen := len(path) - 1; path[pathLen:] == PathSeparator {
		path = path[:pathLen]
	}
	return &Dir{Path: replace(path)}
}

// IsDirectoryExist check directory is exist
func IsExist(path string) bool {
	return false
}

// Permission show file or directory permission
func Permission(fullname string) string {
	return ""
}

// IsFileExist returns true if file exist
func IsFileExist(fullname string) bool {
	return false
}

// Create files or directories
func Create(path string, overwrite bool) error {
	return errors.New("not implement yet")
}

// List returns directory all files or directories
// If nest is true will show all nest files and directories
func List(path string, nest bool) (int, []string, error) {
	return 0, []string{}, errors.New("not implement yet")
}

// Move fils or directories
func Move(dest string, src ...string) error {
	return errors.New("not implement yet")
}

// Copy files or directories
func Copy(dest string, src ...string) error {
	return errors.New("not implement yet")
}

// Delete files or directories
func Delete(fullname ...string) error {
	return errors.New("not implement yet")
}

// replace mismatch separator to right os separator
func replace(path string) string {
	if path == "" {
		return path
	}

	if strings.Contains(path, windowsSeparator) {
		path = strings.ReplaceAll(path, windowsSeparator, PathSeparator)
	} else {
		path = strings.ReplaceAll(path, separator, PathSeparator)
	}
	return path
}
