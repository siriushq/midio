// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package cmd

import (
	"context"

	"github.com/minio/minio-go/v7/pkg/tags"
	"github.com/siriushq/midio/pkg/bucket/policy"
	"github.com/siriushq/midio/pkg/madmin"
)

type ObjectLayerBase struct{}

func (a ObjectLayerBase) BackendInfo() madmin.BackendInfo {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) LocalStorageInfo(ctx context.Context) (StorageInfo, []error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) NSScanner(ctx context.Context, bf *bloomFilter, updates chan<- madmin.DataUsageInfo) error {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) PutObjectMetadata(ctx context.Context, bucket, object string, opts ObjectOptions) (ObjectInfo, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) NewNSLock(bucket string, objects ...string) RWLocker {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) SetDriveCounts() []int {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) ListMultipartUploads(ctx context.Context, bucket string, prefix string, keyMarker string, uploadIDMarker string, delimiter string, maxUploads int) (lmi ListMultipartsInfo, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) NewMultipartUpload(ctx context.Context, bucket string, object string, opts ObjectOptions) (uploadID string, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) CopyObjectPart(ctx context.Context, srcBucket, srcObject, destBucket, destObject, uploadID string, partID int, startOffset, length int64, srcInfo ObjectInfo, srcOpts, dstOpts ObjectOptions) (pi PartInfo, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) PutObjectPart(ctx context.Context, bucket string, object string, uploadID string, partID int, data *PutObjReader, opts ObjectOptions) (pi PartInfo, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) GetMultipartInfo(ctx context.Context, bucket string, object string, uploadID string, opts ObjectOptions) (MultipartInfo, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) ListObjectVersions(ctx context.Context, bucket, prefix, marker, versionMarker, delimiter string, maxKeys int) (ListObjectVersionsInfo, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) ListObjectParts(ctx context.Context, bucket string, object string, uploadID string, partNumberMarker int, maxParts int, opts ObjectOptions) (lpi ListPartsInfo, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) AbortMultipartUpload(ctx context.Context, bucket string, object string, uploadID string, opts ObjectOptions) error {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) CompleteMultipartUpload(ctx context.Context, bucket string, object string, uploadID string, uploadedParts []CompletePart, opts ObjectOptions) (oi ObjectInfo, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) SetBucketPolicy(ctx context.Context, bucket string, bucketPolicy *policy.Policy) error {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) GetBucketPolicy(ctx context.Context, bucket string) (bucketPolicy *policy.Policy, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) DeleteBucketPolicy(ctx context.Context, bucket string) error {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) HealFormat(ctx context.Context, dryRun bool) (madmin.HealResultItem, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) HealBucket(ctx context.Context, bucket string, opts madmin.HealOpts) (madmin.HealResultItem, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) HealObject(ctx context.Context, bucket, object, versionID string, opts madmin.HealOpts) (h madmin.HealResultItem, e error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) ListObjectsV2(ctx context.Context, bucket, prefix, continuationToken, delimiter string, maxKeys int, fetchOwner bool, startAfter string) (result ListObjectsV2Info, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) Walk(ctx context.Context, bucket, prefix string, results chan<- ObjectInfo, opts ObjectOptions) error {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) HealObjects(ctx context.Context, bucket, prefix string, opts madmin.HealOpts, fn HealObjectFn) (e error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) CopyObject(ctx context.Context, srcBucket string, srcObject string, destBucket string, destObject string,
	srcInfo ObjectInfo, srcOpts, dstOpts ObjectOptions) (objInfo ObjectInfo, err error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) GetMetrics(ctx context.Context) (*BackendMetrics, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) PutObjectTags(ctx context.Context, bucket, object string, tags string, opts ObjectOptions) (ObjectInfo, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) GetObjectTags(ctx context.Context, bucket, object string, opts ObjectOptions) (*tags.Tags, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) DeleteObjectTags(ctx context.Context, bucket, object string, opts ObjectOptions) (ObjectInfo, error) {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) IsNotificationSupported() bool {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) IsListenSupported() bool {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) IsEncryptionSupported() bool {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) IsTaggingSupported() bool {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) IsCompressionSupported() bool {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) Health(ctx context.Context, opts HealthOptions) HealthResult {
	panic("function not implemented by this ObjectLayer")
}

func (a ObjectLayerBase) ReadHealth(ctx context.Context) bool {
	panic("function not implemented by this ObjectLayer")
}
