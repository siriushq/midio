package cmd

import (
	"testing"
)

// Tests - canonicalizeETag()
func TestCanonicalizeETag(t *testing.T) {
	testCases := []struct {
		etag              string
		canonicalizedETag string
	}{
		{
			etag:              "\"\"\"",
			canonicalizedETag: "",
		},
		{
			etag:              "\"\"\"abc\"",
			canonicalizedETag: "abc",
		},
		{
			etag:              "abcd",
			canonicalizedETag: "abcd",
		},
		{
			etag:              "abcd\"\"",
			canonicalizedETag: "abcd",
		},
	}
	for _, test := range testCases {
		etag := canonicalizeETag(test.etag)
		if test.canonicalizedETag != etag {
			t.Fatalf("Expected %s , got %s", test.canonicalizedETag, etag)

		}
	}
}
