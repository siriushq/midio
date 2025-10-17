package s3

import (
	minio "github.com/siriushq/midio/cmd"
)

// List of header keys to be filtered, usually
// from all S3 API http responses.
var defaultFilterKeys = []string{
	"Connection",
	"Transfer-Encoding",
	"Accept-Ranges",
	"Date",
	"Server",
	"Vary",
	"x-amz-bucket-region",
	"x-amz-request-id",
	"x-amz-id-2",
	"Content-Security-Policy",
	"X-Xss-Protection",

	// Add new headers to be ignored.
}

// FromGatewayObjectPart converts ObjectInfo for custom part stored as object to PartInfo
func FromGatewayObjectPart(partID int, oi minio.ObjectInfo) (pi minio.PartInfo) {
	return minio.PartInfo{
		Size:         oi.Size,
		ETag:         minio.CanonicalizeETag(oi.ETag),
		LastModified: oi.ModTime,
		PartNumber:   partID,
	}
}
