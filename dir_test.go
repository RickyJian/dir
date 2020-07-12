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
}

func TestCreate(t *testing.T) {
}

func TestDelete(t *testing.T) {
}

func TestIsFileExist(t *testing.T) {
}

func TestList(t *testing.T) {
}

func TestMove(t *testing.T) {
}

func TestPermission(t *testing.T) {
}

func TestReplace(t *testing.T) {
	var tests = []*struct {
		path     string
		expected string
	}{
		{
			path:     "/1/2/3/4",
			expected: "/1/2/3/4",
		},
		{
			path:     "\\1\\2\\3\\4",
			expected: "/1/2/3/4",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, replace(test.path))
	}
}
