package cmd

import (
	"os"
	"testing"
)

// Tests - mkdirAll()
func TestOSMkdirAll(t *testing.T) {
	// create xlStorage test setup
	_, path, err := newXLStorageTestSetup()
	if err != nil {
		t.Fatalf("Unable to create xlStorage test setup, %s", err)
	}
	defer os.RemoveAll(path)

	if err = mkdirAll("", 0777); err != errInvalidArgument {
		t.Fatal("Unexpected error", err)
	}

	if err = mkdirAll(pathJoin(path, "my-obj-del-0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"), 0777); err != errFileNameTooLong {
		t.Fatal("Unexpected error", err)
	}

	if err = mkdirAll(pathJoin(path, "success-vol", "success-object"), 0777); err != nil {
		t.Fatal("Unexpected error", err)
	}
}

// Tests - renameAll()
func TestOSRenameAll(t *testing.T) {
	// create xlStorage test setup
	_, path, err := newXLStorageTestSetup()
	if err != nil {
		t.Fatalf("Unable to create xlStorage test setup, %s", err)
	}
	defer os.RemoveAll(path)

	if err = mkdirAll(pathJoin(path, "testvolume1"), 0777); err != nil {
		t.Fatal(err)
	}
	if err = renameAll("", "foo"); err != errInvalidArgument {
		t.Fatal(err)
	}
	if err = renameAll("foo", ""); err != errInvalidArgument {
		t.Fatal(err)
	}
	if err = renameAll(pathJoin(path, "testvolume1"), pathJoin(path, "testvolume2")); err != nil {
		t.Fatal(err)
	}
	if err = renameAll(pathJoin(path, "testvolume1"), pathJoin(path, "testvolume2")); err != errFileNotFound {
		t.Fatal(err)
	}
	if err = renameAll(pathJoin(path, "my-obj-del-0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"), pathJoin(path, "testvolume2")); err != errFileNameTooLong {
		t.Fatal("Unexpected error", err)
	}
	if err = renameAll(pathJoin(path, "testvolume1"), pathJoin(path, "my-obj-del-0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001")); err != errFileNameTooLong {
		t.Fatal("Unexpected error", err)
	}
}
