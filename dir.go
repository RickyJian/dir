package dir

import (
	"errors"
)

// Dir defines basic directory struct
type Dir struct {
	path       string
	nodes      []string
	fileName   string
	files      []string
	permission string
}

// New new dir
func New(path string) *Dir {
	return new(Dir)
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
