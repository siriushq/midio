package gateway

import (
	// Import all gateways please keep the order

	// NAS
	_ "github.com/siriushq/midio/cmd/gateway/nas"

	// Azure
	_ "github.com/siriushq/midio/cmd/gateway/azure"

	// S3
	_ "github.com/siriushq/midio/cmd/gateway/s3"

	// HDFS
	_ "github.com/siriushq/midio/cmd/gateway/hdfs"

	// GCS (use only if you must, GCS already supports S3 API)
	_ "github.com/siriushq/midio/cmd/gateway/gcs"
	// gateway functionality is frozen, no new gateways are being implemented
	// or considered for upstream inclusion at this point in time. if needed
	// please keep a fork of the project.
)
