//go:build !windows
// +build !windows

package cmd

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/djherbis/atime"
)

// Return error if Atime is disabled on the O/S
func checkAtimeSupport(dir string) (err error) {
	file, err := ioutil.TempFile(dir, "prefix")
	if err != nil {
		return
	}
	defer os.Remove(file.Name())
	defer file.Close()
	finfo1, err := os.Stat(file.Name())
	if err != nil {
		return
	}
	// add a sleep to ensure atime change is detected
	time.Sleep(10 * time.Millisecond)

	if _, err = io.Copy(ioutil.Discard, file); err != nil {
		return
	}

	finfo2, err := os.Stat(file.Name())

	if atime.Get(finfo2).Equal(atime.Get(finfo1)) {
		return errors.New("Atime not supported")
	}
	return
}
