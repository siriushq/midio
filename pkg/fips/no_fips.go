package fips

import (
	"crypto/tls"

	"github.com/minio/sio"
)

var enabled = false

func cipherSuitesDARE() []byte {
	return []byte{sio.AES_256_GCM, sio.CHACHA20_POLY1305}
}

func cipherSuitesTLS() []uint16 {
	return []uint16{
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	}
}

func ellipticCurvesTLS() []tls.CurveID {
	return []tls.CurveID{tls.X25519, tls.CurveP256}
}
