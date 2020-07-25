package dir

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMode(t *testing.T) {
	var tests = []*struct {
		permission os.FileMode
		expected   *Mode
	}{
		{
			permission: 0000,
			expected: &Mode{
				User: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
				Group: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
				Other: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
			},
		},
		{
			permission: 0001,
			expected: &Mode{
				User: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
				Group: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
				Other: &info{
					Read:  false,
					Write: false,
					Exec:  true,
				},
			},
		},
		{
			permission: 0040,
			expected: &Mode{
				User: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
				Group: &info{
					Read:  true,
					Write: false,
					Exec:  false,
				},
				Other: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
			},
		},
		{
			permission: 0700,
			expected: &Mode{
				User: &info{
					Read:  true,
					Write: true,
					Exec:  true,
				},
				Group: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
				Other: &info{
					Read:  false,
					Write: false,
					Exec:  false,
				},
			},
		},
		{
			permission: 0552,
			expected: &Mode{
				User: &info{
					Read:  true,
					Write: false,
					Exec:  true,
				},
				Group: &info{
					Read:  true,
					Write: false,
					Exec:  true,
				},
				Other: &info{
					Read:  false,
					Write: true,
					Exec:  false,
				},
			},
		},
		{
			permission: os.ModePerm,
			expected: &Mode{
				User: &info{
					Read:  true,
					Write: true,
					Exec:  true,
				},
				Group: &info{
					Read:  true,
					Write: true,
					Exec:  true,
				},
				Other: &info{
					Read:  true,
					Write: true,
					Exec:  true,
				},
			},
		},
	}
	for _, test := range tests {
		mode := ParseMode(test.permission)
		assert.Equal(t, mode.User.Read, mode.User.Read)
		assert.Equal(t, mode.User.Write, mode.User.Write)
		assert.Equal(t, mode.User.Exec, mode.User.Exec)
		assert.Equal(t, mode.Group.Read, mode.Group.Read)
		assert.Equal(t, mode.Group.Write, mode.Group.Write)
		assert.Equal(t, mode.Group.Exec, mode.Group.Exec)
		assert.Equal(t, mode.Other.Read, mode.Other.Read)
		assert.Equal(t, mode.Other.Write, mode.Other.Write)
		assert.Equal(t, mode.Other.Exec, mode.Other.Exec)
	}
}

// TODO: random test
