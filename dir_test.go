package dir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var tests = []*struct {
		dir      *Dir
		expected *Dir
	}{
		{
			dir: New("/tmp/stickers/"),
			expected: &Dir{
				path:     "/tmp/stickers",
				nodes:    []string{"tmp", "stickers"},
				fileName: "",
			},
		},
		{
			dir: New("/tmp/stickers"),
			expected: &Dir{
				path:     "/tmp/stickers",
				nodes:    []string{"tmp", "stickers"},
				fileName: "",
			},
		},
		{
			dir: New("/tmp/stickers/dogs"),
			expected: &Dir{
				path:     "/tmp/stickers/dogs",
				nodes:    []string{"tmp", "stickers", "dogs"},
				fileName: "",
			},
		},
		{
			dir: New("/tmp/stickers/sticker.jpg"),
			expected: &Dir{
				path:     "/tmp/stickers",
				nodes:    []string{"tmp", "stickers", "dogs"},
				fileName: "sticker.jpg",
			},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected.path, test.dir.path)
		assert.Equal(t, test.expected.nodes, test.dir.nodes)
		assert.Equal(t, test.expected.fileName, test.dir.fileName)
	}
}

func TestCopy(t *testing.T) {
}

func TestCreate(t *testing.T) {
}

func TestDelete(t *testing.T) {
}

func TestIsExist(t *testing.T) {
}

func TestIsFileExist(t *testing.T) {
}

func TestList(t *testing.T) {
}

func TestMove(t *testing.T) {
}

func TestPermission(t *testing.T) {
}
