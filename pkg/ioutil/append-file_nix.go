//go:build !windows
// +build !windows

package ioutil

import (
	"io"
	"os"
)

// AppendFile - appends the file "src" to the file "dst"
func AppendFile(dst string, src string, osync bool) error {
	flags := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	if osync {
		flags = flags | os.O_SYNC
	}
	appendFile, err := os.OpenFile(dst, flags, 0666)
	if err != nil {
		return err
	}
	defer appendFile.Close()

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	_, err = io.Copy(appendFile, srcFile)
	return err
}
