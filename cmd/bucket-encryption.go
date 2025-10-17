package cmd

import (
	"errors"
	"io"

	bucketsse "github.com/siriushq/midio/pkg/bucket/encryption"
)

// BucketSSEConfigSys - in-memory cache of bucket encryption config
type BucketSSEConfigSys struct{}

// NewBucketSSEConfigSys - Creates an empty in-memory bucket encryption configuration cache
func NewBucketSSEConfigSys() *BucketSSEConfigSys {
	return &BucketSSEConfigSys{}
}

// Get - gets bucket encryption config for the given bucket.
func (sys *BucketSSEConfigSys) Get(bucket string) (*bucketsse.BucketSSEConfig, error) {
	if globalIsGateway {
		objAPI := newObjectLayerFn()
		if objAPI == nil {
			return nil, errServerNotInitialized
		}

		return nil, BucketSSEConfigNotFound{Bucket: bucket}
	}

	return globalBucketMetadataSys.GetSSEConfig(bucket)
}

// validateBucketSSEConfig parses bucket encryption configuration and validates if it is supported by MinIO.
func validateBucketSSEConfig(r io.Reader) (*bucketsse.BucketSSEConfig, error) {
	encConfig, err := bucketsse.ParseBucketSSEConfig(r)
	if err != nil {
		return nil, err
	}

	if len(encConfig.Rules) == 1 && encConfig.Rules[0].DefaultEncryptionAction.Algorithm == bucketsse.AES256 {
		return encConfig, nil
	}

	return nil, errors.New("Unsupported bucket encryption configuration")
}
