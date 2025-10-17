package cmd

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/siriushq/midio/cmd/crypto"
	"github.com/siriushq/midio/pkg/hash"
)

var toAPIErrorTests = []struct {
	err     error
	errCode APIErrorCode
}{
	{err: hash.BadDigest{}, errCode: ErrBadDigest},
	{err: hash.SHA256Mismatch{}, errCode: ErrContentSHA256Mismatch},
	{err: IncompleteBody{}, errCode: ErrIncompleteBody},
	{err: ObjectExistsAsDirectory{}, errCode: ErrObjectExistsAsDirectory},
	{err: BucketNameInvalid{}, errCode: ErrInvalidBucketName},
	{err: BucketExists{}, errCode: ErrBucketAlreadyOwnedByYou},
	{err: ObjectNotFound{}, errCode: ErrNoSuchKey},
	{err: ObjectNameInvalid{}, errCode: ErrInvalidObjectName},
	{err: InvalidUploadID{}, errCode: ErrNoSuchUpload},
	{err: InvalidPart{}, errCode: ErrInvalidPart},
	{err: InsufficientReadQuorum{}, errCode: ErrSlowDown},
	{err: InsufficientWriteQuorum{}, errCode: ErrSlowDown},
	{err: InvalidMarkerPrefixCombination{}, errCode: ErrNotImplemented},
	{err: InvalidUploadIDKeyCombination{}, errCode: ErrNotImplemented},
	{err: MalformedUploadID{}, errCode: ErrNoSuchUpload},
	{err: PartTooSmall{}, errCode: ErrEntityTooSmall},
	{err: BucketNotEmpty{}, errCode: ErrBucketNotEmpty},
	{err: BucketNotFound{}, errCode: ErrNoSuchBucket},
	{err: StorageFull{}, errCode: ErrStorageFull},
	{err: NotImplemented{}, errCode: ErrNotImplemented},
	{err: errSignatureMismatch, errCode: ErrSignatureDoesNotMatch},

	// SSE-C errors
	{err: crypto.ErrInvalidCustomerAlgorithm, errCode: ErrInvalidSSECustomerAlgorithm},
	{err: crypto.ErrMissingCustomerKey, errCode: ErrMissingSSECustomerKey},
	{err: crypto.ErrInvalidCustomerKey, errCode: ErrAccessDenied},
	{err: crypto.ErrMissingCustomerKeyMD5, errCode: ErrMissingSSECustomerKeyMD5},
	{err: crypto.ErrCustomerKeyMD5Mismatch, errCode: ErrSSECustomerKeyMD5Mismatch},
	{err: errObjectTampered, errCode: ErrObjectTampered},

	{err: nil, errCode: ErrNone},
	{err: errors.New("Custom error"), errCode: ErrInternalError}, // Case where err type is unknown.
}

func TestAPIErrCode(t *testing.T) {
	disk := filepath.Join(globalTestTmpDir, "minio-"+nextSuffix())
	defer os.RemoveAll(disk)

	initFSObjects(disk, t)

	ctx := context.Background()
	for i, testCase := range toAPIErrorTests {
		errCode := toAPIErrorCode(ctx, testCase.err)
		if errCode != testCase.errCode {
			t.Errorf("Test %d: Expected error code %d, got %d", i+1, testCase.errCode, errCode)
		}
	}
}
