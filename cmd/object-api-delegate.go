// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package cmd

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7/pkg/tags"
	"github.com/siriushq/midio/pkg/bucket/policy"
	"github.com/siriushq/midio/pkg/madmin"
)

type ObjectLayerDelegate struct{}

// BackendInfo returns the underlying backend information
func (a ObjectLayerDelegate) BackendInfo() madmin.BackendInfo {
	if globalObjectAPI == nil {
		return madmin.BackendInfo{}
	}

	return globalObjectAPI.BackendInfo()
}

// LocalStorageInfo returns the local disks information, mainly used
// in prometheus - for gateway this just a no-op
func (a ObjectLayerDelegate) LocalStorageInfo(ctx context.Context) (StorageInfo, []error) {
	if globalObjectAPI == nil {
		return StorageInfo{}, []error{errors.New("object layer not initialized")}
	}

	return globalObjectAPI.LocalStorageInfo(ctx)
}

// NSScanner - scanner is not implemented for gateway
func (a ObjectLayerDelegate) NSScanner(ctx context.Context, bf *bloomFilter, updates chan<- madmin.DataUsageInfo) error {
	if globalObjectAPI == nil {
		return errors.New("object layer not initialized")
	}

	return globalObjectAPI.NSScanner(ctx, bf, updates)
}

// PutObjectMetadata - not implemented for gateway.
func (a ObjectLayerDelegate) PutObjectMetadata(ctx context.Context, bucket, object string, opts ObjectOptions) (ObjectInfo, error) {
	if globalObjectAPI == nil {
		return ObjectInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.PutObjectMetadata(ctx, bucket, object, opts)
}

// NewNSLock is a dummy stub for gateway.
func (a ObjectLayerDelegate) NewNSLock(bucket string, objects ...string) RWLocker {
	if globalObjectAPI == nil {
		return nil
	}

	return globalObjectAPI.NewNSLock(bucket, objects...)
}

// SetDriveCounts no-op
func (a ObjectLayerDelegate) SetDriveCounts() []int {
	if globalObjectAPI == nil {
		return nil
	}

	return globalObjectAPI.SetDriveCounts()
}

// ListMultipartUploads lists all multipart uploads.
func (a ObjectLayerDelegate) ListMultipartUploads(ctx context.Context, bucket string, prefix string, keyMarker string, uploadIDMarker string, delimiter string, maxUploads int) (lmi ListMultipartsInfo, err error) {
	if globalObjectAPI == nil {
		return ListMultipartsInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.ListMultipartUploads(ctx, bucket, prefix, keyMarker, uploadIDMarker, delimiter, maxUploads)
}

// NewMultipartUpload upload object in multiple parts
func (a ObjectLayerDelegate) NewMultipartUpload(ctx context.Context, bucket string, object string, opts ObjectOptions) (uploadID string, err error) {
	if globalObjectAPI == nil {
		return "", errors.New("object layer not initialized")
	}

	return globalObjectAPI.NewMultipartUpload(ctx, bucket, object, opts)
}

// CopyObjectPart copy part of object to uploadID for another object
func (a ObjectLayerDelegate) CopyObjectPart(ctx context.Context, srcBucket, srcObject, destBucket, destObject, uploadID string, partID int, startOffset, length int64, srcInfo ObjectInfo, srcOpts, dstOpts ObjectOptions) (pi PartInfo, err error) {
	if globalObjectAPI == nil {
		return PartInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.CopyObjectPart(ctx, srcBucket, srcObject, destBucket, destObject, uploadID, partID, startOffset, length, srcInfo, srcOpts, dstOpts)
}

// PutObjectPart puts a part of object in bucket
func (a ObjectLayerDelegate) PutObjectPart(ctx context.Context, bucket string, object string, uploadID string, partID int, data *PutObjReader, opts ObjectOptions) (pi PartInfo, err error) {
	if globalObjectAPI == nil {
		return PartInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.PutObjectPart(ctx, bucket, object, uploadID, partID, data, opts)
}

// GetMultipartInfo returns metadata associated with the uploadId
func (a ObjectLayerDelegate) GetMultipartInfo(ctx context.Context, bucket string, object string, uploadID string, opts ObjectOptions) (MultipartInfo, error) {
	if globalObjectAPI == nil {
		return MultipartInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.GetMultipartInfo(ctx, bucket, object, uploadID, opts)
}

// ListObjectVersions returns all object parts for specified object in specified bucket
func (a ObjectLayerDelegate) ListObjectVersions(ctx context.Context, bucket, prefix, marker, versionMarker, delimiter string, maxKeys int) (ListObjectVersionsInfo, error) {
	if globalObjectAPI == nil {
		return ListObjectVersionsInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.ListObjectVersions(ctx, bucket, prefix, marker, versionMarker, delimiter, maxKeys)
}

// ListObjectParts returns all object parts for specified object in specified bucket
func (a ObjectLayerDelegate) ListObjectParts(ctx context.Context, bucket string, object string, uploadID string, partNumberMarker int, maxParts int, opts ObjectOptions) (lpi ListPartsInfo, err error) {
	if globalObjectAPI == nil {
		return ListPartsInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.ListObjectParts(ctx, bucket, object, uploadID, partNumberMarker, maxParts, opts)
}

// AbortMultipartUpload aborts a ongoing multipart upload
func (a ObjectLayerDelegate) AbortMultipartUpload(ctx context.Context, bucket string, object string, uploadID string, opts ObjectOptions) error {
	if globalObjectAPI == nil {
		return errors.New("object layer not initialized")
	}

	return globalObjectAPI.AbortMultipartUpload(ctx, bucket, object, uploadID, opts)
}

// CompleteMultipartUpload completes ongoing multipart upload and finalizes object
func (a ObjectLayerDelegate) CompleteMultipartUpload(ctx context.Context, bucket string, object string, uploadID string, uploadedParts []CompletePart, opts ObjectOptions) (oi ObjectInfo, err error) {
	if globalObjectAPI == nil {
		return ObjectInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.CompleteMultipartUpload(ctx, bucket, object, uploadID, uploadedParts, opts)
}

// SetBucketPolicy sets policy on bucket
func (a ObjectLayerDelegate) SetBucketPolicy(ctx context.Context, bucket string, bucketPolicy *policy.Policy) error {
	if globalObjectAPI == nil {
		return errors.New("object layer not initialized")
	}

	return globalObjectAPI.SetBucketPolicy(ctx, bucket, bucketPolicy)
}

// GetBucketPolicy will get policy on bucket
func (a ObjectLayerDelegate) GetBucketPolicy(ctx context.Context, bucket string) (bucketPolicy *policy.Policy, err error) {
	if globalObjectAPI == nil {
		return nil, errors.New("object layer not initialized")
	}

	return globalObjectAPI.GetBucketPolicy(ctx, bucket)
}

// DeleteBucketPolicy deletes all policies on bucket
func (a ObjectLayerDelegate) DeleteBucketPolicy(ctx context.Context, bucket string) error {
	if globalObjectAPI == nil {
		return errors.New("object layer not initialized")
	}

	return globalObjectAPI.DeleteBucketPolicy(ctx, bucket)
}

//// SetBucketVersioning enables versioning on a bucket.
//func (a ObjectLayerDelegate) SetBucketVersioning(ctx context.Context, bucket string, v *versioning.Versioning) error {
//	return globalObjectAPI.SetBucketVersioning(ctx, bucket, v)
//}
//
//// GetBucketVersioning retrieves versioning configuration of a bucket.
//func (a ObjectLayerDelegate) GetBucketVersioning(ctx context.Context, bucket string) (*versioning.Versioning, error) {
//	return globalObjectAPI.GetBucketVersioning(ctx, bucket)
//}
//
//// SetBucketLifecycle enables lifecycle policies on a bucket.
//func (a ObjectLayerDelegate) SetBucketLifecycle(ctx context.Context, bucket string, lifecycle *lifecycle.Lifecycle) error {
//	return globalObjectAPI.SetBucketLifecycle(ctx, bucket, lifecycle)
//}
//
//// GetBucketLifecycle retrieves lifecycle configuration of a bucket.
//func (a ObjectLayerDelegate) GetBucketLifecycle(ctx context.Context, bucket string) (*lifecycle.Lifecycle, error) {
//	return globalObjectAPI.GetBucketLifecycle(ctx, bucket)
//}
//
//// DeleteBucketLifecycle deletes all lifecycle policies on a bucket
//func (a ObjectLayerDelegate) DeleteBucketLifecycle(ctx context.Context, bucket string) error {
//	return globalObjectAPI.DeleteBucketLifecycle(ctx, bucket)
//}
//
//// GetBucketSSEConfig returns bucket encryption config on a bucket
//func (a ObjectLayerDelegate) GetBucketSSEConfig(ctx context.Context, bucket string) (*bucketsse.BucketSSEConfig, error) {
//	return globalObjectAPI.GetBucketSSEConfig(ctx, bucket)
//}
//
//// SetBucketSSEConfig sets bucket encryption config on a bucket
//func (a ObjectLayerDelegate) SetBucketSSEConfig(ctx context.Context, bucket string, config *bucketsse.BucketSSEConfig) error {
//	return globalObjectAPI.SetBucketSSEConfig(ctx, bucket, config)
//}
//
//// DeleteBucketSSEConfig deletes bucket encryption config on a bucket
//func (a ObjectLayerDelegate) DeleteBucketSSEConfig(ctx context.Context, bucket string) error {
//	return globalObjectAPI.DeleteBucketSSEConfig(ctx, bucket)
//}

// HealFormat - Not implemented stub
func (a ObjectLayerDelegate) HealFormat(ctx context.Context, dryRun bool) (madmin.HealResultItem, error) {
	if globalObjectAPI == nil {
		return madmin.HealResultItem{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.HealFormat(ctx, dryRun)
}

// HealBucket - Not implemented stub
func (a ObjectLayerDelegate) HealBucket(ctx context.Context, bucket string, opts madmin.HealOpts) (madmin.HealResultItem, error) {
	if globalObjectAPI == nil {
		return madmin.HealResultItem{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.HealBucket(ctx, bucket, opts)
}

// HealObject - Not implemented stub
func (a ObjectLayerDelegate) HealObject(ctx context.Context, bucket, object, versionID string, opts madmin.HealOpts) (h madmin.HealResultItem, e error) {
	if globalObjectAPI == nil {
		return madmin.HealResultItem{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.HealObject(ctx, bucket, object, versionID, opts)
}

// ListObjectsV2 - Not implemented stub
func (a ObjectLayerDelegate) ListObjectsV2(ctx context.Context, bucket, prefix, continuationToken, delimiter string, maxKeys int, fetchOwner bool, startAfter string) (result ListObjectsV2Info, err error) {
	if globalObjectAPI == nil {
		return ListObjectsV2Info{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.ListObjectsV2(ctx, bucket, prefix, continuationToken, delimiter, maxKeys, fetchOwner, startAfter)
}

// Walk - Not implemented stub
func (a ObjectLayerDelegate) Walk(ctx context.Context, bucket, prefix string, results chan<- ObjectInfo, opts ObjectOptions) error {
	if globalObjectAPI == nil {
		return errors.New("object layer not initialized")
	}

	return globalObjectAPI.Walk(ctx, bucket, prefix, results, opts)
}

// HealObjects - Not implemented stub
func (a ObjectLayerDelegate) HealObjects(ctx context.Context, bucket, prefix string, opts madmin.HealOpts, fn HealObjectFn) (e error) {
	if globalObjectAPI == nil {
		return errors.New("object layer not initialized")
	}

	return globalObjectAPI.HealObjects(ctx, bucket, prefix, opts, fn)
}

// CopyObject copies a blob from source container to destination container.
func (a ObjectLayerDelegate) CopyObject(ctx context.Context, srcBucket string, srcObject string, destBucket string, destObject string,
	srcInfo ObjectInfo, srcOpts, dstOpts ObjectOptions) (objInfo ObjectInfo, err error) {
	if globalObjectAPI == nil {
		return ObjectInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.CopyObject(ctx, srcBucket, srcObject, destBucket, destObject, srcInfo, srcOpts, dstOpts)
}

// GetMetrics - no op
func (a ObjectLayerDelegate) GetMetrics(ctx context.Context) (*BackendMetrics, error) {
	if globalObjectAPI == nil {
		return nil, errors.New("object layer not initialized")
	}

	return globalObjectAPI.GetMetrics(ctx)
}

// PutObjectTags - not implemented.
func (a ObjectLayerDelegate) PutObjectTags(ctx context.Context, bucket, object string, tags string, opts ObjectOptions) (ObjectInfo, error) {
	if globalObjectAPI == nil {
		return ObjectInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.PutObjectTags(ctx, bucket, object, tags, opts)
}

// GetObjectTags - not implemented.
func (a ObjectLayerDelegate) GetObjectTags(ctx context.Context, bucket, object string, opts ObjectOptions) (*tags.Tags, error) {
	if globalObjectAPI == nil {
		return nil, errors.New("object layer not initialized")
	}

	return globalObjectAPI.GetObjectTags(ctx, bucket, object, opts)
}

// DeleteObjectTags - not implemented.
func (a ObjectLayerDelegate) DeleteObjectTags(ctx context.Context, bucket, object string, opts ObjectOptions) (ObjectInfo, error) {
	if globalObjectAPI == nil {
		return ObjectInfo{}, errors.New("object layer not initialized")
	}

	return globalObjectAPI.DeleteObjectTags(ctx, bucket, object, opts)
}

// IsNotificationSupported returns whether bucket notification is applicable for this layer.
func (a ObjectLayerDelegate) IsNotificationSupported() bool {
	return globalObjectAPI != nil && globalObjectAPI.IsNotificationSupported()
}

// IsListenSupported returns whether listen bucket notification is applicable for this layer.
func (a ObjectLayerDelegate) IsListenSupported() bool {
	return globalObjectAPI != nil && globalObjectAPI.IsListenSupported()
}

// IsEncryptionSupported returns whether server side encryption is implemented for this layer.
func (a ObjectLayerDelegate) IsEncryptionSupported() bool {
	return globalObjectAPI != nil && globalObjectAPI.IsEncryptionSupported()
}

// IsTaggingSupported returns whether object tagging is supported or not for this layer.
func (a ObjectLayerDelegate) IsTaggingSupported() bool {
	return globalObjectAPI != nil && globalObjectAPI.IsTaggingSupported()
}

// IsCompressionSupported returns whether compression is applicable for this layer.
func (a ObjectLayerDelegate) IsCompressionSupported() bool {
	return globalObjectAPI != nil && globalObjectAPI.IsCompressionSupported()
}

// Health - No Op.
func (a ObjectLayerDelegate) Health(ctx context.Context, opts HealthOptions) HealthResult {
	if globalObjectAPI == nil {
		return HealthResult{}
	}

	return globalObjectAPI.Health(ctx, opts)
}

// ReadHealth - No Op.
func (a ObjectLayerDelegate) ReadHealth(ctx context.Context) bool {
	return globalObjectAPI != nil && globalObjectAPI.ReadHealth(ctx)
}
