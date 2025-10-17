package cmd

import (
	"reflect"
	"testing"
)

// Tests cache exclude parsing.
func TestParseGatewaySSE(t *testing.T) {
	testCases := []struct {
		gwSSEStr string
		expected gatewaySSE
		success  bool
	}{
		// valid input
		{"c;S3", []string{"C", "S3"}, true},
		{"S3", []string{"S3"}, true},
		{"c,S3", []string{}, false},
		{"c;S3;KMS", []string{}, false},
		{"C;s3", []string{"C", "S3"}, true},
	}

	for i, testCase := range testCases {
		gwSSE, err := parseGatewaySSE(testCase.gwSSEStr)
		if err != nil && testCase.success {
			t.Errorf("Test %d: Expected success but failed instead %s", i+1, err)
		}
		if err == nil && !testCase.success {
			t.Errorf("Test %d: Expected failure but passed instead", i+1)
		}
		if err == nil {
			if !reflect.DeepEqual(gwSSE, testCase.expected) {
				t.Errorf("Test %d: Expected %v, got %v", i+1, testCase.expected, gwSSE)
			}
		}
	}
}
