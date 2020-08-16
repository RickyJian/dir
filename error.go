package dir

import "errors"

var (
	// ErrFileOrDirectoryNotExist defines file or directory not exist
	ErrFileOrDirectoryNotExist = errors.New("file or directory not found")
	// ErrPathIsNotDirectory defines path is not directory
	ErrPathIsNotDirectory = errors.New("path is not directory")
	// ErrEmptyPath defines path is empty
	ErrEmptyPath = errors.New("empty path")
	// ErrCannotDeleteRoot defines cannot delete root
	ErrCannotDeleteRoot = errors.New("cannot delete root")
)
