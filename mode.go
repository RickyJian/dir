package dir

import (
	"os"
)

const (
	userShift  = 6
	groupShift = 3
	read       = 4
	write      = 2
	exec       = 1
)

// Mode defines user group other permission
type Mode struct {
	User  *info
	Group *info
	Other *info
}

// info defines basic linux permission
type info struct {
	Read  bool
	Write bool
	Exec  bool
}

// ParseMode parse os.FileMode to mode
func ParseMode(mode os.FileMode) *Mode {
	return &Mode{
		User: &info{
			Read:  mode.Perm()&(read<<userShift) != 0,
			Write: mode.Perm()&(write<<userShift) != 0,
			Exec:  mode.Perm()&(exec<<userShift) != 0,
		},
		Group: &info{
			Read:  mode.Perm()&(read<<groupShift) != 0,
			Write: mode.Perm()&(write<<groupShift) != 0,
			Exec:  mode.Perm()&(exec<<groupShift) != 0,
		},
		Other: &info{
			Read:  mode.Perm()&read != 0,
			Write: mode.Perm()&write != 0,
			Exec:  mode.Perm()&exec != 0,
		},
	}
}
