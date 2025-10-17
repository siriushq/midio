package policy

import (
	"testing"
)

func TestEffectIsAllowed(t *testing.T) {
	testCases := []struct {
		effect         Effect
		check          bool
		expectedResult bool
	}{
		{Allow, false, false},
		{Allow, true, true},
		{Deny, false, true},
		{Deny, true, false},
	}

	for i, testCase := range testCases {
		result := testCase.effect.IsAllowed(testCase.check)

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}

}

func TestEffectIsValid(t *testing.T) {
	testCases := []struct {
		effect         Effect
		expectedResult bool
	}{
		{Allow, true},
		{Deny, true},
		{Effect(""), false},
		{Effect("foo"), false},
	}

	for i, testCase := range testCases {
		result := testCase.effect.IsValid()

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}
