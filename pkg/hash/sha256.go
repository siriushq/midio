package hash

import (
	"crypto/sha256"
	"hash"
)

// newSHA256 returns a new hash.Hash computing the SHA256 checksum.
// The SHA256 implementation is FIPS 140-2 compliant when the
// boringcrypto branch of Go is used.
// Ref: https://github.com/golang/go/tree/dev.boringcrypto
func newSHA256() hash.Hash { return sha256.New() }
