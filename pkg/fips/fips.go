//go:build fips
// +build fips

package fips

import (
	"crypto/tls"

	"github.com/minio/sio"
)

var enabled = true

func cipherSuitesDARE() []byte {
	return []byte{sio.AES_256_GCM}
}

func cipherSuitesTLS() []uint16 {
	return []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	}
}

func ellipticCurvesTLS() []tls.CurveID {
	return []tls.CurveID{tls.CurveP256}
}
