package sql

import (
	"testing"
)

func TestEvalSQLLike(t *testing.T) {
	dropCases := []struct {
		input, resultExpected string
		matchExpected         bool
	}{
		{"", "", false},
		{"a", "", true},
		{"ab", "b", true},
		{"தமிழ்", "மிழ்", true},
	}

	for i, tc := range dropCases {
		res, ok := dropRune(tc.input)
		if res != tc.resultExpected || ok != tc.matchExpected {
			t.Errorf("DropRune Case %d failed", i)
		}
	}

	matcherCases := []struct {
		iText, iPat        string
		iHasLeadingPercent bool
		resultExpected     string
		matchExpected      bool
	}{
		{"abcd", "bcd", false, "", false},
		{"abcd", "bcd", true, "", true},
		{"abcd", "abcd", false, "", true},
		{"abcd", "abcd", true, "", true},
		{"abcd", "ab", false, "cd", true},
		{"abcd", "ab", true, "cd", true},
		{"abcd", "bc", false, "", false},
		{"abcd", "bc", true, "d", true},
	}

	for i, tc := range matcherCases {
		res, ok := matcher(tc.iText, tc.iPat, tc.iHasLeadingPercent)
		if res != tc.resultExpected || ok != tc.matchExpected {
			t.Errorf("Matcher Case %d failed", i)
		}
	}

	evalCases := []struct {
		iText, iPat   string
		iEsc          rune
		matchExpected bool
		errExpected   error
	}{
		{"abcd", "abc", runeZero, false, nil},
		{"abcd", "abcd", runeZero, true, nil},
		{"abcd", "abc_", runeZero, true, nil},
		{"abcd", "_bdd", runeZero, false, nil},
		{"abcd", "_b_d", runeZero, true, nil},

		{"abcd", "____", runeZero, true, nil},
		{"abcd", "____%", runeZero, true, nil},
		{"abcd", "%____", runeZero, true, nil},
		{"abcd", "%__", runeZero, true, nil},
		{"abcd", "____", runeZero, true, nil},

		{"", "_", runeZero, false, nil},
		{"", "%", runeZero, true, nil},
		{"abcd", "%%%%%", runeZero, true, nil},
		{"abcd", "_____", runeZero, false, nil},
		{"abcd", "%%%%%", runeZero, true, nil},

		{"a%%d", `a\%\%d`, '\\', true, nil},
		{"a%%d", `a\%d`, '\\', false, nil},
		{`a%%\d`, `a\%\%\\d`, '\\', true, nil},
		{`a%%\`, `a\%\%\\`, '\\', true, nil},
		{`a%__%\`, `a\%\_\_\%\\`, '\\', true, nil},

		{`a%__%\`, `a\%\_\_\%_`, '\\', true, nil},
		{`a%__%\`, `a\%\_\__`, '\\', false, nil},
		{`a%__%\`, `a\%\_\_%`, '\\', true, nil},
		{`a%__%\`, `a?%?_?_?%\`, '?', true, nil},
	}

	for i, tc := range evalCases {
		// fmt.Println("Case:", i)
		res, err := evalSQLLike(tc.iText, tc.iPat, tc.iEsc)
		if res != tc.matchExpected || err != tc.errExpected {
			t.Errorf("Eval Case %d failed: %v %v", i, res, err)
		}
	}
}

func TestEvalSQLSubstring(t *testing.T) {
	evalCases := []struct {
		s           string
		startIdx    int
		length      int
		resExpected string
		errExpected error
	}{
		{"abcd", 1, 1, "a", nil},
		{"abcd", -1, 1, "a", nil},
		{"abcd", 999, 999, "", nil},
		{"", 999, 999, "", nil},
		{"测试abc", 1, 1, "测", nil},
		{"测试abc", 5, 5, "c", nil},
	}

	for i, tc := range evalCases {
		res, err := evalSQLSubstring(tc.s, tc.startIdx, tc.length)
		if res != tc.resExpected || err != tc.errExpected {
			t.Errorf("Eval Case %d failed: %v %v", i, res, err)
		}
	}
}
