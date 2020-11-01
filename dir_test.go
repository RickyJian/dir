package dir

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var tests = []*struct {
		dir      *Dir
		expected *Dir
	}{
		{
			dir:      New("/tmp/stickers/"),
			expected: &Dir{Path: "/tmp/stickers"},
		},
		{
			dir:      New("\\tmp\\stickers\\"),
			expected: &Dir{Path: "/tmp/stickers"},
		},
		{
			dir:      New("/tmp/stickers"),
			expected: &Dir{Path: "/tmp/stickers"},
		},
		{
			dir:      New("\\tmp\\stickers"),
			expected: &Dir{Path: "/tmp/stickers"},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected.Path, test.dir.Path)
	}
}

func TestIsExist(t *testing.T) {
	var tests = []*struct {
		path         string
		expected     bool
		expectedFile os.FileInfo
	}{
		{
			path:     "/",
			expected: true,
		},
		{
			path:     "dir.go",
			expected: true,
		},
		{
			path:     "dir2.go",
			expected: false,
		},
	}
	for _, test := range tests {
		f, ok := IsExist(test.path)
		assert.Equal(t, test.expected, ok)
		if !ok {
			assert.Nil(t, f)
		}
	}
}

func TestCopy(t *testing.T) {
	// TODO: mock test
}

func TestCreate(t *testing.T) {
	// TODO: mock test
}

func TestList(t *testing.T) {
	var tests = []*struct {
		path          string
		hidden        bool
		expectedNo    int
		expectedError error
	}{
		{
			path:          "dir1",
			expectedNo:    -1,
			expectedError: ErrDirectoryNotFound,
		},
		{
			path:          "../dir",
			hidden:        false,
			expectedNo:    6,
			expectedError: nil,
		},
		{
			path:          "dir.go",
			expectedNo:    -1,
			expectedError: ErrPathIsNotDirectory,
		},
	}
	for _, test := range tests {
		no, _, err := List(test.path, test.hidden)
		assert.Equal(t, test.expectedNo, no)
		assert.Equal(t, test.expectedError, err)
	}
}

func TestMove(t *testing.T) {
	// TODO: mock test
}

func TestDelete(t *testing.T) {
	// TODO: mock test
}

func TestIsValid(t *testing.T) {
	var tests = []*struct {
		op       Operation
		expected bool
	}{
		{
			op:       Default,
			expected: true,
		},
		{
			op:       Merge,
			expected: true,
		},
		{
			op:       Override,
			expected: true,
		},
		{
			op:       5,
			expected: false,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, test.op.isValid())
	}
}
