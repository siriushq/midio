package cmd

import (
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	Version = "2017-05-07T06:37:49Z"
	_, err := time.Parse(time.RFC3339, Version)
	if err != nil {
		t.Fatal(err)
	}
}
