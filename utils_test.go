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
