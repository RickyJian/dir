package dir

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	nodes := strings.Split(filepath.Dir(path), PathSeparator)
	return &Dir{
		Path:  replace(path),
		Nodes: nodes,
	}
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
func (d *Dir) List(hidden bool) (int, error) {
	no, files, err := List(d.Path, hidden)
	if err != nil {
		return -1, err
	}
	d.Files = files
	return no, nil
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
func List(route string, hidden bool) (int, []string, error) {
	f, ok := IsExist(route)
	if !ok {
		return -1, []string{}, ErrFileOrDirectoryNotExist
	} else if !f.IsDir() {
		return -1, []string{}, ErrPathIsNotDirectory
	}

	routeLen := len(route)
	if route[routeLen-1:] != PathSeparator {
		routeLen++
	}
	var files []string
	err := filepath.Walk(route, func(path string, info os.FileInfo, err error) error {
		if route == path {
			// skip first directory
			return nil
		} else if file := filepath.Base(path); !hidden && file[:1] == "." {
			return filepath.SkipDir
		}
		files = append(files, path[routeLen:])
		return nil
	})
	if err != nil {
		return -1, []string{}, fmt.Errorf("walk through directory failed: %w", err)
	}
	return len(files), files, nil
}

// Move files or directories
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
