package policy

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/minio/minio-go/v7/pkg/set"
)

func TestPrincipalIsValid(t *testing.T) {
	testCases := []struct {
		principal      Principal
		expectedResult bool
	}{
		{NewPrincipal("*"), true},
		{NewPrincipal("arn:aws:iam::AccountNumber:root"), true},
		{NewPrincipal(), false},
	}

	for i, testCase := range testCases {
		result := testCase.principal.IsValid()

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestPrincipalIntersection(t *testing.T) {
	testCases := []struct {
		principal            Principal
		principalToIntersect Principal
		expectedResult       set.StringSet
	}{
		{NewPrincipal("*"), NewPrincipal("*"), set.CreateStringSet("*")},
		{NewPrincipal("arn:aws:iam::AccountNumber:root"), NewPrincipal("arn:aws:iam::AccountNumber:myuser"), set.CreateStringSet()},
		{NewPrincipal(), NewPrincipal("*"), set.CreateStringSet()},
	}

	for i, testCase := range testCases {
		result := testCase.principal.Intersection(testCase.principalToIntersect)

		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestPrincipalMarshalJSON(t *testing.T) {
	testCases := []struct {
		principal      Principal
		expectedResult []byte
		expectErr      bool
	}{
		{NewPrincipal("*"), []byte(`{"AWS":["*"]}`), false},
		{NewPrincipal("arn:aws:iam::AccountNumber:*"), []byte(`{"AWS":["arn:aws:iam::AccountNumber:*"]}`), false},
		{NewPrincipal(), nil, true},
	}

	for i, testCase := range testCases {
		result, err := json.Marshal(testCase.principal)
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, string(testCase.expectedResult), string(result))
			}
		}
	}
}

func TestPrincipalMatch(t *testing.T) {
	testCases := []struct {
		principals     Principal
		principal      string
		expectedResult bool
	}{
		{NewPrincipal("*"), "AccountNumber", true},
		{NewPrincipal("arn:aws:iam::*"), "arn:aws:iam::AccountNumber:root", true},
		{NewPrincipal("arn:aws:iam::AccountNumber:*"), "arn:aws:iam::TestAccountNumber:root", false},
	}

	for i, testCase := range testCases {
		result := testCase.principals.Match(testCase.principal)

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestPrincipalUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		data           []byte
		expectedResult Principal
		expectErr      bool
	}{
		{[]byte(`"*"`), NewPrincipal("*"), false},
		{[]byte(`{"AWS": "*"}`), NewPrincipal("*"), false},
		{[]byte(`{"AWS": "arn:aws:iam::AccountNumber:*"}`), NewPrincipal("arn:aws:iam::AccountNumber:*"), false},
		{[]byte(`"arn:aws:iam::AccountNumber:*"`), NewPrincipal(), true},
		{[]byte(`["arn:aws:iam::AccountNumber:*", "arn:aws:iam:AnotherAccount:*"]`), NewPrincipal(), true},
	}

	for i, testCase := range testCases {
		var result Principal
		err := json.Unmarshal(testCase.data, &result)
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
			}
		}
	}
}
