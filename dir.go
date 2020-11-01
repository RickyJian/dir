package dir

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	Files      []string
	Subs       []string
}

// New new dir
func New(path string) *Dir {
	path = Replace(path)
	if pathLen := len(path) - 1; path[pathLen:] == PathSeparator {
		path = path[:pathLen]
	}
	return &Dir{Path: path}
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

// List subdirectories and files
func (d *Dir) List(hidden bool) error {
	sub, files, err := List(d.Path, hidden)
	if err != nil {
		return err
	}
	d.Files = files
	d.Subs = sub
	return nil
}

// Operation defines dir operation
type Operation int32

const (
	// Default:
	//   * move - if destination directory exist, it will not move directory and return `ErrDestExist`.
	Default Operation = iota
	// Merge:
	//   * move - if destination directory exist, it will move src files which are not exist in destination directories,
	//            and if files are exist it will override it.
	Merge
	// Override:
	//   * move - if destination directory exist, it will override it.
	Override
)

var moveOperationSet = map[Operation]struct{}{
	Default:  {},
	Merge:    {},
	Override: {},
}

// isValid check move operation valid
func (o Operation) isValid() bool {
	_, ok := moveOperationSet[o]
	return ok
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

// List subdirectories and files
func List(route string, hidden bool) ([]string, []string, error) {
	routeLen := len(route)
	if route[routeLen-1:] != PathSeparator {
		routeLen++
	}

	var files, subs []string
	err := filepath.Walk(route, func(path string, info os.FileInfo, err error) error {
		if route == path {
			// skip first directory
			return nil
		} else if file := filepath.Base(path); !hidden && file[:1] == "." {
			return filepath.SkipDir
		} else if isDir := info.IsDir(); isDir {
			subs = append(subs, path[routeLen:])
		} else if !isDir {
			files = append(files, path[routeLen:])
		}
		return nil
	})
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("walk through directory failed: %w", err)
	}
	return subs, files, nil
}

// Copy files or directories
func Copy(dest string, src ...string) error {
	return errors.New("not implement yet")
}
