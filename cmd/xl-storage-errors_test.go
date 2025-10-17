package cmd

import (
	"os"
	"runtime"
	"syscall"
	"testing"
)

func TestSysErrors(t *testing.T) {
	pathErr := &os.PathError{Err: syscall.ENAMETOOLONG}
	ok := isSysErrTooLong(pathErr)
	if !ok {
		t.Fatalf("Unexpected error expecting %s", syscall.ENAMETOOLONG)
	}
	pathErr = &os.PathError{Err: syscall.ENOTDIR}
	ok = isSysErrNotDir(pathErr)
	if !ok {
		t.Fatalf("Unexpected error expecting %s", syscall.ENOTDIR)
	}
	if runtime.GOOS != globalWindowsOSName {
		pathErr = &os.PathError{Err: syscall.ENOTEMPTY}
		ok = isSysErrNotEmpty(pathErr)
		if !ok {
			t.Fatalf("Unexpected error expecting %s", syscall.ENOTEMPTY)
		}
	} else {
		pathErr = &os.PathError{Err: syscall.Errno(0x91)}
		ok = isSysErrNotEmpty(pathErr)
		if !ok {
			t.Fatal("Unexpected error expecting 0x91")
		}
	}
	if runtime.GOOS == globalWindowsOSName {
		pathErr = &os.PathError{Err: syscall.Errno(0x03)}
		ok = isSysErrPathNotFound(pathErr)
		if !ok {
			t.Fatal("Unexpected error expecting 0x03")
		}
	}
}
