package dir

import "errors"

var (
	// ErrFileOrDirectoryNotExist defines file or directory not exist
	ErrFileOrDirectoryNotExist = errors.New("file or directory not found")
	// ErrPathIsNotDirectory defines path is not directory
	ErrPathIsNotDirectory = errors.New("path is not directory")
	// ErrEmptyPath defines path is empty
	ErrEmptyPath = errors.New("empty path")
	// ErrEmptySrc defines src is empty
	ErrEmptySrc = errors.New("empty src")
	// ErrEmptyDest defines dest is empty
	ErrEmptyDest = errors.New("empty dest")
	// ErrCannotDeleteRoot defines cannot delete root
	ErrCannotDeleteRoot = errors.New("cannot delete root")
	// ErrOperationInvalid operation is invalid
	ErrOperationInvalid = errors.New("operation invalid")
)
