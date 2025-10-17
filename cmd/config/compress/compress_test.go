package compress

import (
	"reflect"
	"testing"
)

func TestParseCompressIncludes(t *testing.T) {
	testCases := []struct {
		str              string
		expectedPatterns []string
		success          bool
	}{
		// invalid input
		{",,,", []string{}, false},
		{"", []string{}, false},
		{",", []string{}, false},
		{"/", []string{}, false},
		{"text/*,/", []string{}, false},

		// valid input
		{".txt,.log", []string{".txt", ".log"}, true},
		{"text/*,application/json", []string{"text/*", "application/json"}, true},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.str, func(t *testing.T) {
			gotPatterns, err := parseCompressIncludes(testCase.str)
			if !testCase.success && err == nil {
				t.Error("expected failure but success instead")
			}
			if testCase.success && err != nil {
				t.Errorf("expected success but failed instead %s", err)
			}
			if testCase.success && !reflect.DeepEqual(testCase.expectedPatterns, gotPatterns) {
				t.Errorf("expected patterns %s but got %s", testCase.expectedPatterns, gotPatterns)
			}
		})
	}
}
