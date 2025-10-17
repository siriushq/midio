//go:build !windows && !darwin && !freebsd
// +build !windows,!darwin,!freebsd

package ioutil

import (
	"os"
)

var readMode = os.O_RDONLY | 0x40000 // read with O_NOATIME
