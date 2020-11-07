package dir

import "errors"

var (
	// ErrDirectoryNotFound defines directory not found
	ErrDirectoryNotFound = errors.New("directory not found")
	// ErrPathIsNotDirectory defines path is not directory
	ErrPathIsNotDirectory = errors.New("path is not directory")
	// ErrDirectoryExist defines directory exist
	ErrDirectoryExist = errors.New("directory not found")
	// ErrEmptyPath defines path is empty
	ErrEmptyPath = errors.New("empty path")
	// ErrCannotDeleteRoot defines cannot delete root
	ErrCannotDeleteRoot = errors.New("cannot delete root")
	// ErrInvalidOperation defines operation is invalid
	ErrInvalidOperation = errors.New("invalid operation")
)
