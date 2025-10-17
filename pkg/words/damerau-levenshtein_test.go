package words

import (
	"math"
	"testing"
)

// Test minimum function which calculates the minimal value in a list of integers
func TestMinimum(t *testing.T) {
	type testCase struct {
		listval  []int
		expected int
	}
	testCases := []testCase{
		{listval: []int{3, 4, 15}, expected: 3},
		{listval: []int{}, expected: math.MaxInt32},
	}
	// Validate all the test cases.
	for i, tt := range testCases {
		val := minimum(tt.listval)
		if val != tt.expected {
			t.Errorf("Test %d:, Expected %d, got %d", i+1, tt.expected, val)
		}
	}
}

// Test DamerauLevenshtein which calculates the difference distance between two words
func TestDamerauLevenshtein(t *testing.T) {
	type testCase struct {
		word1    string
		word2    string
		distance int
	}
	testCases := []testCase{
		{word1: "", word2: "", distance: 0},
		{word1: "a", word2: "a", distance: 0},
		{word1: "a", word2: "b", distance: 1},
		{word1: "rm", word2: "tm", distance: 1},
		{word1: "version", word2: "evrsion", distance: 1},
		{word1: "version", word2: "bersio", distance: 2},
	}
	// Validate all the test cases.
	for i, tt := range testCases {
		d := DamerauLevenshteinDistance(tt.word1, tt.word2)
		if d != tt.distance {
			t.Errorf("Test %d:, Expected %d, got %d", i+1, tt.distance, d)
		}
	}
}
