//go:build linux
// +build linux

package cmd

import "testing"

// Tests allocate.
func TestFallocate(t *testing.T) {
	err := Fallocate(0, 0, 0)
	if err != nil {
		t.Fatal("Unexpected error in fallocate for length 0:", err)
	}
}
