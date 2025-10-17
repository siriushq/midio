//go:build fips
// +build fips

package cmd

var (
	// Newer official download info URLs appear earlier below.
	minioReleaseInfoURL = minioReleaseURL + "minio.fips.sha256sum"
)
