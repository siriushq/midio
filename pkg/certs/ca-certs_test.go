package certs

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetRootCAs(t *testing.T) {
	emptydir, err := ioutil.TempDir("", "test-get-root-cas")
	if err != nil {
		t.Fatalf("Unable create temp directory. %v", emptydir)
	}
	defer os.RemoveAll(emptydir)

	dir1, err := ioutil.TempDir("", "test-get-root-cas")
	if err != nil {
		t.Fatalf("Unable create temp directory. %v", dir1)
	}
	defer os.RemoveAll(dir1)
	if err = os.Mkdir(filepath.Join(dir1, "empty-dir"), 0755); err != nil {
		t.Fatalf("Unable create empty dir. %v", err)
	}

	dir2, err := ioutil.TempDir("", "test-get-root-cas")
	if err != nil {
		t.Fatalf("Unable create temp directory. %v", dir2)
	}
	defer os.RemoveAll(dir2)
	if err = ioutil.WriteFile(filepath.Join(dir2, "empty-file"), []byte{}, 0644); err != nil {
		t.Fatalf("Unable create test file. %v", err)
	}

	testCases := []struct {
		certCAsDir  string
		expectedErr error
	}{
		// ignores non-existent directories.
		{"nonexistent-dir", nil},
		// Ignores directories.
		{dir1, nil},
		// Ignore empty directory.
		{emptydir, nil},
		// Loads the cert properly.
		{dir2, nil},
	}

	for _, testCase := range testCases {
		_, err := GetRootCAs(testCase.certCAsDir)

		if testCase.expectedErr == nil {
			if err != nil {
				t.Fatalf("error: expected = <nil>, got = %v", err)
			}
		} else if err == nil {
			t.Fatalf("error: expected = %v, got = <nil>", testCase.expectedErr)
		} else if testCase.expectedErr.Error() != err.Error() {
			t.Fatalf("error: expected = %v, got = %v", testCase.expectedErr, err)
		}
	}
}
