package dir

import (
	"errors"
	"fmt"
	"io/ioutil"
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
	Files      []string
	Subs       []string
}

// New new dir
func New(path string) *Dir {
	path = Replace(path)
	if pathLen := len(path) - 1; path[pathLen:] == PathSeparator {
		path = path[:pathLen]
	}
	return &Dir{
		Path: path,
		Name: filepath.Base(path),
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

// Get current path directories and files
func (d *Dir) Get(hidden bool) ([]string, []string, error) {
	return Get(d.Path, hidden)
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

// Move directory(include files)
func (d *Dir) Move(dest string, op Operation) error {
	return Move(dest, d, op)
}

// Delete directory(include files)
func (d *Dir) Delete() error {
	return Delete(d.Path)
}

// IsExist check path is exist and return os.fileInfo
func IsExist(path string) (os.FileInfo, bool) {
	file, err := os.Stat(path)
	if err != nil {
		return nil, false
	}
	return file, true
}

// Get current path directories and files
func Get(path string, hidden bool) ([]string, []string, error) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("failed to read directory: %w", err)
	}

	var dirs, files []string
	for _, info := range infos {
		if name := info.Name(); info.IsDir() {
			dirs = append(dirs, name)
		} else {
			if !hidden && name[:1] == Hidden {
				continue
			}
			files = append(files, name)
		}
	}
	return dirs, files, nil
}

// List subdirectories and files
func List(path string, hidden bool) ([]string, []string, error) {
	routeLen := len(path)
	if path[routeLen-1:] != PathSeparator {
		routeLen++
	}

	var subs, files []string
	err := filepath.Walk(path, func(route string, info os.FileInfo, err error) error {
		if path == route {
			// skip first directory
			return nil
		} else if !hidden && info.Name()[:1] == Hidden {
			return filepath.SkipDir
		} else if info.IsDir() {
			subs = append(subs, route[routeLen:])
		} else {
			files = append(files, route[routeLen:])
		}
		return nil
	})
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("walk through directory failed: %w", err)
	}
	return subs, files, nil
}

// Move directories(include files)
func Move(dest string, src *Dir, op Operation) error {
	if !op.isValid() {
		return ErrInvalidOperation
	} else if dest = strings.TrimSpace(dest); dest == "" {
		return ErrEmptyPath
	}
	name := filepath.Base(dest)
	if filepath.Dir(dest) != name {
		return ErrInvalidPath
	}

	dest = Replace(dest)
	switch op {
	case Default:
		return move(dest, src)
	case Merge:
		return moveMerge(dest, src)
	case Override:
		return moveOverride(dest, src)
	}
	src.Path = dest
	src.Name = name
	return nil
}

// move directories(include files) which operation is `Default`
func move(dest string, src *Dir) error {
	if _, ok := IsExist(dest); ok {
		return ErrDirectoryExist
	} else if err := os.Rename(src.Path, dest); err != nil {
		return fmt.Errorf("failed to move directory: %w", err)
	}
	return nil
}

func moveMerge(dest string, src *Dir) error {
	return nil
}

func moveOverride(dest string, src *Dir) error {
	if _, ok := IsExist(dest); ok {
		if err := deleteDir(dest); err != nil {
			return err
		}
	}

	if err := os.Rename(src.Path, dest); err != nil {
		return fmt.Errorf("failed to move directory: %w", err)
	}
	return nil
}

// Delete directory(include files)
func Delete(path string) error {
	if path = strings.TrimSpace(path); path == "" {
		return ErrEmptyPath
	} else if filepath.Dir(path) != filepath.Base(path) {
		return ErrInvalidPath
	} else if path == Root {
		return ErrCannotDeleteRoot
	}
	return deleteDir(path)
}

func deleteDir(path string) error {
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

// Copy files or directories
func Copy(dest string, src ...string) error {
	return errors.New("not implement yet")
}
