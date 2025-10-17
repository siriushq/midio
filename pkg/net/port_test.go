package net

import (
	"testing"
)

func TestPortString(t *testing.T) {
	testCases := []struct {
		port        Port
		expectedStr string
	}{
		{Port(0), "0"},
		{Port(9000), "9000"},
		{Port(65535), "65535"},
		{Port(1024), "1024"},
	}

	for i, testCase := range testCases {
		str := testCase.port.String()

		if str != testCase.expectedStr {
			t.Fatalf("test %v: error: port: %v, got: %v", i+1, testCase.expectedStr, str)
		}
	}
}

func TestParsePort(t *testing.T) {
	testCases := []struct {
		s            string
		expectedPort Port
		expectErr    bool
	}{
		{"0", Port(0), false},
		{"9000", Port(9000), false},
		{"65535", Port(65535), false},
		{"http", Port(80), false},
		{"https", Port(443), false},
		{"90000", Port(0), true},
		{"-10", Port(0), true},
		{"", Port(0), true},
		{" 1024", Port(0), true},
	}

	for i, testCase := range testCases {
		port, err := ParsePort(testCase.s)
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("test %v: error: expected: %v, got: %v", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if port != testCase.expectedPort {
				t.Fatalf("test %v: error: port: %v, got: %v", i+1, testCase.expectedPort, port)
			}
		}
	}
}
