package iampolicy

import (
	"testing"
)

func TestActionIsObjectAction(t *testing.T) {
	testCases := []struct {
		action         Action
		expectedResult bool
	}{
		{AbortMultipartUploadAction, true},
		{DeleteObjectAction, true},
		{GetObjectAction, true},
		{ListMultipartUploadPartsAction, true},
		{PutObjectAction, true},
		{CreateBucketAction, false},
	}

	for i, testCase := range testCases {
		result := testCase.action.isObjectAction()

		if testCase.expectedResult != result {
			t.Fatalf("case %v: expected: %v, got: %v", i+1, testCase.expectedResult, result)
		}
	}
}

func TestActionIsValid(t *testing.T) {
	testCases := []struct {
		action         Action
		expectedResult bool
	}{
		{PutObjectAction, true},
		{AbortMultipartUploadAction, true},
		{Action("foo"), false},
	}

	for i, testCase := range testCases {
		result := testCase.action.IsValid()

		if testCase.expectedResult != result {
			t.Fatalf("case %v: expected: %v, got: %v", i+1, testCase.expectedResult, result)
		}
	}
}
