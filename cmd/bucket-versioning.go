package cmd

import "github.com/siriushq/midio/pkg/bucket/versioning"

// BucketVersioningSys - policy subsystem.
type BucketVersioningSys struct{}

// Enabled enabled versioning?
func (sys *BucketVersioningSys) Enabled(bucket string) bool {
	vc, err := globalBucketMetadataSys.GetVersioningConfig(bucket)
	if err != nil {
		return false
	}
	return vc.Enabled()
}

// Suspended suspended versioning?
func (sys *BucketVersioningSys) Suspended(bucket string) bool {
	vc, err := globalBucketMetadataSys.GetVersioningConfig(bucket)
	if err != nil {
		return false
	}
	return vc.Suspended()
}

// Get returns stored bucket policy
func (sys *BucketVersioningSys) Get(bucket string) (*versioning.Versioning, error) {
	if globalIsGateway {
		objAPI := newObjectLayerFn()
		if objAPI == nil {
			return nil, errServerNotInitialized
		}
		return nil, NotImplemented{}
	}
	return globalBucketMetadataSys.GetVersioningConfig(bucket)
}

// Reset BucketVersioningSys to initial state.
func (sys *BucketVersioningSys) Reset() {
	// There is currently no internal state.
}

// NewBucketVersioningSys - creates new versioning system.
func NewBucketVersioningSys() *BucketVersioningSys {
	return &BucketVersioningSys{}
}
