package policy

import (
	"testing"
)

func TestIDIsValid(t *testing.T) {
	testCases := []struct {
		id             ID
		expectedResult bool
	}{
		{ID("DenyEncryptionSt1"), true},
		{ID(""), true},
		{ID("aa\xe2"), false},
	}

	for i, testCase := range testCases {
		result := testCase.id.IsValid()

		if result != testCase.expectedResult {
			t.Errorf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}
