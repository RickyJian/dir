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
	File       os.FileInfo
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

// IsExist check path exist and set File
func (d *Dir) IsExist() bool {
	file, ok := IsExist(d.Path)
	if !ok {
		return false
	}
	d.File = file
	return true
}

// List returns directory all files or directories
// If nest is true will show all nest files or directories
func (d *Dir) List(nest bool) (int, []string, error) {
	no, nodes, err := List(d.Path, nest)
	if err != nil {
		return -1, []string{}, err
	}
	d.Nodes = nodes
	return no, nodes, nil
}

// IsExist check path is exist and return os.fileInfo
func IsExist(path string) (os.FileInfo, bool) {
	file, err := os.Stat(path)
	if err != nil {
		return nil, false
	}
	return file, true
}

// Permission show file or directory permission
func Permission(fullname string) string {
	return ""
}

// Create files or directories
func Create(path string, overwrite bool) error {
	return errors.New("not implement yet")
}

// List returns directory all files or directories
// If nest is true will show all nest files or directories
func List(path string, nest bool) (int, []string, error) {
	// TODO: nest
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
