package cmd

import (
	"testing"
	"time"
)

// Tests ToObjectInfo function.
func TestCacheMetadataObjInfo(t *testing.T) {
	m := cacheMeta{Meta: nil}
	objInfo := m.ToObjectInfo("testbucket", "testobject")
	if objInfo.Size != 0 {
		t.Fatal("Unexpected object info value for Size", objInfo.Size)
	}
	if !objInfo.ModTime.Equal(time.Time{}) {
		t.Fatal("Unexpected object info value for ModTime ", objInfo.ModTime)
	}
	if objInfo.IsDir {
		t.Fatal("Unexpected object info value for IsDir", objInfo.IsDir)
	}
	if !objInfo.Expires.IsZero() {
		t.Fatal("Unexpected object info value for Expires ", objInfo.Expires)
	}
}

// test wildcard patterns for excluding entries from cache
func TestCacheExclusion(t *testing.T) {
	cobjects := &cacheObjects{
		cache: nil,
	}

	testCases := []struct {
		bucketName     string
		objectName     string
		excludePattern string
		expectedResult bool
	}{
		{"testbucket", "testobjectmatch", "testbucket/testobj*", true},
		{"testbucket", "testobjectnomatch", "testbucet/testobject*", false},
		{"testbucket", "testobject/pref1/obj1", "*/*", true},
		{"testbucket", "testobject/pref1/obj1", "*/pref1/*", true},
		{"testbucket", "testobject/pref1/obj1", "testobject/*", false},
		{"photos", "image1.jpg", "*.jpg", true},
		{"photos", "europe/paris/seine.jpg", "seine.jpg", false},
		{"photos", "europe/paris/seine.jpg", "*/seine.jpg", true},
		{"phil", "z/likes/coffee", "*/likes/*", true},
		{"failbucket", "no/slash/prefixes", "/failbucket/no/", false},
		{"failbucket", "no/slash/prefixes", "/failbucket/no/*", false},
	}

	for i, testCase := range testCases {
		cobjects.exclude = []string{testCase.excludePattern}
		if cobjects.isCacheExclude(testCase.bucketName, testCase.objectName) != testCase.expectedResult {
			t.Fatal("Cache exclusion test failed for case ", i)
		}
	}
}
