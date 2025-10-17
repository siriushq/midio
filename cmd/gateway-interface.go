package cmd

import (
	"github.com/siriushq/midio/pkg/auth"
)

// GatewayMinioSysTmp prefix is used in Azure/GCS gateway for save metadata sent by Initialize Multipart Upload API.
const (
	GatewayMinioSysTmp  = "minio.sys.tmp/"
	AzureBackendGateway = "azure"
	GCSBackendGateway   = "gcs"
	HDFSBackendGateway  = "hdfs"
	NASBackendGateway   = "nas"
	S3BackendGateway    = "s3"
)

// Gateway represents a gateway backend.
type Gateway interface {
	// Name returns the unique name of the gateway.
	Name() string

	// NewGatewayLayer returns a new  ObjectLayer.
	NewGatewayLayer(creds auth.Credentials) (ObjectLayer, error)

	// Returns true if gateway is ready for production.
	Production() bool
}
