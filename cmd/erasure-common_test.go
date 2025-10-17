package cmd

import (
	"bytes"
	"context"
	"os"
	"testing"
)

// Tests for if parent directory is object
func TestErasureParentDirIsObject(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	obj, fsDisks, err := prepareErasureSets32(ctx)
	if err != nil {
		t.Fatalf("Unable to initialize 'Erasure' object layer.")
	}
	defer obj.Shutdown(context.Background())

	// Remove all disks.
	for _, disk := range fsDisks {
		defer os.RemoveAll(disk)
	}

	bucketName := "testbucket"
	objectName := "object"

	if err = obj.MakeBucketWithLocation(GlobalContext, bucketName, BucketOptions{}); err != nil {
		t.Fatal(err)
	}

	objectContent := "12345"
	_, err = obj.PutObject(GlobalContext, bucketName, objectName,
		mustGetPutObjReader(t, bytes.NewReader([]byte(objectContent)), int64(len(objectContent)), "", ""), ObjectOptions{})
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		expectedErr bool
		objectName  string
	}{
		{
			expectedErr: true,
			objectName:  pathJoin(objectName, "parent-is-object"),
		},
		{
			expectedErr: false,
			objectName:  pathJoin("no-parent", "object"),
		},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			_, err = obj.PutObject(GlobalContext, bucketName, testCase.objectName,
				mustGetPutObjReader(t, bytes.NewReader([]byte(objectContent)), int64(len(objectContent)), "", ""), ObjectOptions{})
			if testCase.expectedErr && err == nil {
				t.Error("Expected error but got nil")
			}
			if !testCase.expectedErr && err != nil {
				t.Errorf("Expected nil but got %v", err)
			}
		})
	}
}
