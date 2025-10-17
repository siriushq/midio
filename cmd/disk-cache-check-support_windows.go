//go:build windows
// +build windows

package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/djherbis/atime"
	"golang.org/x/sys/windows/registry"
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
	atime.Get(finfo1)

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\FileSystem`, registry.QUERY_VALUE)
	if err != nil {
		return
	}
	defer k.Close()

	setting, _, err := k.GetIntegerValue("NtfsDisableLastAccessUpdate")
	if err != nil {
		return
	}

	lowSetting := setting & 0xFFFF
	if lowSetting != uint64(0x0000) && lowSetting != uint64(0x0002) {
		return errors.New("Atime not supported")
	}
	return
}
