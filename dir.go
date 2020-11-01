package dir

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var (
	// PathSeparator defines string type path separator
	PathSeparator = string(os.PathSeparator)
)

// Dir defines basic directory struct
type Dir struct {
	Path       string
	Name       string
	Size       int64
	ModTime    time.Time
	Permission *Mode
	Nodes      []string
	Files      []string
}

// New new dir
func New(path string) *Dir {
	path = Replace(path)
	if pathLen := len(path) - 1; path[pathLen:] == PathSeparator {
		path = path[:pathLen]
	}
	return &Dir{
		Path:  path,
		Nodes: strings.Split(filepath.Dir(path), PathSeparator),
	}
}

// IsExist check path exist and set File
func (d *Dir) IsExist() error {
	file, ok := IsExist(d.Path)
	if !ok {
		return ErrDirectoryNotFound
	} else if !file.IsDir() {
		return ErrPathIsNotDirectory
	}

	d.Name = file.Name()
	d.Size = file.Size()
	d.ModTime = file.ModTime()
	d.Permission = ParseMode(file.Mode())
	return nil
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

// MoveOperation defines move operation
type MoveOperation int32

const (
	// None when destination directory exist, it will return `ErrFileOrDirectoryExist`
	None MoveOperation = iota
	// Merge is directory operation. When destination directory exist, it will copy all files in it.
	Merge
	// Override. When destination directory exist, it will delete source directory then create it.
	Override
)

// Move directory
// When name is not empty it will move file; otherwise, it will move directory.
func (d *Dir) Move(dest string, op MoveOperation) error {
	if dest == "" {
		return ErrEmptyDest
	} else if !isMoveOperationValid(op) {
		return ErrOperationInvalid
	}
	destFile, exist := IsExist(Replace(dest))
	if exist && !destFile.IsDir() && op == Merge {
		return ErrOperationInvalid
	}

	switch op {
	case None:
		return os.Rename(filepath.Join(d.Path, d.Name), dest)
	case Merge:

	case Override:
	}
	return nil
}

// Delete directory
// When Name is empty it will delete directory
func (d *Dir) Delete() error {
	path := d.Path
	if path == "" {
		return ErrEmptyPath
	} else if path == Root {
		return ErrCannotDeleteRoot
	} else if d.Name != "" {
		path = filepath.Join(d.Path, d.Name)
	}

	if err := os.RemoveAll(path); err == nil {
		// do nothing
	} else if pathErr := new(os.PathError); errors.As(err, &pathErr) &&
		!errors.Is(pathErr, syscall.ENOENT) {
		// syscall.ENOENT is `no such file or directory`
		// in this package we will not return this error.
		// if you want to check file exist, you can use
		// `IsExist` function.
		return pathErr.Err
	}
	return nil
}

// IsExist check path is exist and return os.fileInfo
func IsExist(path string) (os.FileInfo, bool) {
	file, err := os.Stat(path)
	if err != nil {
		return nil, false
	}
	return file, true
}

// Create files or directories
func Create(path string, overwrite bool) error {
	return errors.New("not implement yet")
}

// List returns directory all files or directories
func List(route string, hidden bool) (int, []string, error) {
	f, ok := IsExist(route)
	if !ok {
		return -1, []string{}, ErrDirectoryNotFound
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
func Move(src, dest string) error {
	if src == "" {
		return ErrEmptySrc
	} else if dest == "" {
		return ErrEmptyDest
	}

	return os.Rename(src, dest)
}

// Copy files or directories
func Copy(dest string, src ...string) error {
	return errors.New("not implement yet")
}

// isMoveOperationValid check move operation valid
func isMoveOperationValid(op MoveOperation) bool {
	switch op {
	case None,
		Merge,
		Override:
		// valid operation
	default:
		return false
	}
	return true
}
