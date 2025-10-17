package sys

import "testing"

// Test get max open file limit.
func TestGetMaxOpenFileLimit(t *testing.T) {
	_, _, err := GetMaxOpenFileLimit()
	if err != nil {
		t.Errorf("expected: nil, got: %v", err)
	}
}

// Test set open file limit
func TestSetMaxOpenFileLimit(t *testing.T) {
	curLimit, maxLimit, err := GetMaxOpenFileLimit()
	if err != nil {
		t.Fatalf("Unable to get max open file limit. %v", err)
	}

	err = SetMaxOpenFileLimit(curLimit, maxLimit)
	if err != nil {
		t.Errorf("expected: nil, got: %v", err)
	}
}
