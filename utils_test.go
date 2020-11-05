package dir

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	godotenv.Overload("test.env")
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
		{
			path:     "~/tmp",
			expected: "/os/home/tmp",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, Replace(test.path))
	}
}

func TestDiff(t *testing.T) {
	var tests = []*struct {
		src           []string
		dest          []string
		expectedPlus  []string
		expectedMinus []string
	}{
		{
			src:           []string{"/a"},
			dest:          []string{},
			expectedPlus:  []string{"/a"},
			expectedMinus: []string{},
		},
		{
			src:           []string{"a", "/a/b"},
			dest:          []string{},
			expectedPlus:  []string{"a", "/a/b"},
			expectedMinus: []string{},
		},
		{
			src:           []string{},
			dest:          []string{"/a"},
			expectedPlus:  []string{},
			expectedMinus: []string{},
		},
		{
			src:           []string{"/a"},
			dest:          []string{"/a"},
			expectedPlus:  []string{},
			expectedMinus: []string{"/a"},
		},
		{
			src:           []string{"/a", "/b"},
			dest:          []string{"/a"},
			expectedPlus:  []string{"/b"},
			expectedMinus: []string{"/a"},
		},
		{
			src:           []string{"/a/b", "/b/a"},
			dest:          []string{"/a/b", "/b/a"},
			expectedPlus:  []string{},
			expectedMinus: []string{"/a/b", "/b/a"},
		},
	}
	for _, test := range tests {
		plus, minus := Diff(test.dest, test.src)
		for i, p := range plus {
			assert.Equal(t, test.expectedPlus[i], p)
		}
		for i, m := range minus {
			assert.Equal(t, test.expectedMinus[i], m)
		}
	}
}
