//go:build windows
// +build windows

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Test if various paths work as expected when converted to UNC form
func TestUNCPaths(t *testing.T) {
	var testCases = []struct {
		objName string
		pass    bool
	}{
		{"/abcdef", true},
		{"/a/b/c/d/e/f/g", true},
		{string(bytes.Repeat([]byte("界"), 85)), true},
		// Each path component must be <= 255 bytes long.
		{string(bytes.Repeat([]byte("界"), 280)), false},
		{`/p/q/r/s/t`, true},
	}
	dir, err := ioutil.TempDir("", "testdisk-")
	if err != nil {
		t.Fatal(err)
	}
	// Cleanup on exit of test
	defer os.RemoveAll(dir)

	// Instantiate posix object to manage a disk
	var fs StorageAPI
	fs, err = newLocalXLStorage(dir)
	if err != nil {
		t.Fatal(err)
	}

	// Create volume to use in conjunction with other StorageAPI's file API(s)
	err = fs.MakeVol(context.Background(), "voldir")
	if err != nil {
		t.Fatal(err)
	}

	for i, test := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			err = fs.AppendFile(context.Background(), "voldir", test.objName, []byte("hello"))
			if err != nil && test.pass {
				t.Error(err)
			} else if err == nil && !test.pass {
				t.Error(err)
			}
			fs.Delete(context.Background(), "voldir", test.objName, false)
		})
	}
}

// Test to validate xlStorage behavior on windows when a non-final path component is a file.
func TestUNCPathENOTDIR(t *testing.T) {
	// Instantiate posix object to manage a disk
	dir, err := ioutil.TempDir("", "testdisk-")
	if err != nil {
		t.Fatal(err)
	}
	// Cleanup on exit of test
	defer os.RemoveAll(dir)

	var fs StorageAPI
	fs, err = newLocalXLStorage(dir)
	if err != nil {
		t.Fatal(err)
	}

	// Create volume to use in conjunction with other StorageAPI's file API(s)
	err = fs.MakeVol(context.Background(), "voldir")
	if err != nil {
		t.Fatal(err)
	}

	err = fs.AppendFile(context.Background(), "voldir", "/file", []byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	// Try to create a file that includes a file in its path components.
	// In *nix, this returns syscall.ENOTDIR while in windows we receive the following error.
	err = fs.AppendFile(context.Background(), "voldir", "/file/obj1", []byte("hello"))
	if err != errFileAccessDenied {
		t.Errorf("expected: %s, got: %s", errFileAccessDenied, err)
	}
}
