package cmd

import (
	"testing"
)

func TestNewRequestID(t *testing.T) {
	// Ensure that it returns an alphanumeric result of length 16.
	var id = mustGetRequestID(UTCNow())

	if len(id) != 16 {
		t.Fail()
	}

	var e rune
	for _, char := range id {
		e = rune(char)

		// Ensure that it is alphanumeric, in this case, between 0-9 and A-Z.
		if !(('0' <= e && e <= '9') || ('A' <= e && e <= 'Z')) {
			t.Fail()
		}
	}
}
