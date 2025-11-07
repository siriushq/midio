//go:build !netbsd && !solaris
// +build !netbsd,!solaris

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package disk_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/siriushq/midio/pkg/disk"
)

func TestFree(t *testing.T) {
	path, err := ioutil.TempDir(os.TempDir(), "minio-")
	defer os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}

	di, err := disk.GetInfo(path)
	if err != nil {
		t.Fatal(err)
	}

	if di.FSType == "UNKNOWN" {
		t.Error("Unexpected FSType", di.FSType)
	}
}
