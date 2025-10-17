package sys

import "testing"

// Test get max memory limit.
func TestGetMaxMemoryLimit(t *testing.T) {
	_, _, err := GetMaxMemoryLimit()
	if err != nil {
		t.Errorf("expected: nil, got: %v", err)
	}
}

// Test set memory limit
func TestSetMaxMemoryLimit(t *testing.T) {
	curLimit, maxLimit, err := GetMaxMemoryLimit()
	if err != nil {
		t.Fatalf("Unable to get max memory limit. %v", err)
	}

	err = SetMaxMemoryLimit(curLimit, maxLimit)
	if err != nil {
		t.Errorf("expected: nil, got: %v", err)
	}
}
