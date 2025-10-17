// Package fips provides functionality to configure cryptographic
// implementations compliant with FIPS 140.
//
// FIPS 140 [1] is a US standard for data processing that specifies
// requirements for cryptographic modules. Software that is "FIPS 140
// compliant" must use approved cryptographic primitives only and that
// are implemented by a FIPS 140 certified cryptographic module.
//
// So, FIPS 140 requires that a certified implementation of e.g. AES
// is used to implement more high-level cryptographic protocols.
// It does not require any specific security criteria for those
// high-level protocols. FIPS 140 focuses only on the implementation
// and usage of the most low-level cryptographic building blocks.
//
// [1]: https://en.wikipedia.org/wiki/FIPS_140
package fips

import "crypto/tls"

// Enabled returns true if and only if FIPS 140-2 support
// is enabled.
//
// FIPS 140-2 requires that only specifc cryptographic
// primitives, like AES or SHA-256, are used and that
// those primitives are implemented by a FIPS 140-2
// certified cryptographic module.
func Enabled() bool { return enabled }

// CipherSuitesDARE returns the supported cipher suites
// for the DARE object encryption.
func CipherSuitesDARE() []byte {
	return cipherSuitesDARE()
}

// CipherSuitesTLS returns the supported cipher suites
// used by the TLS stack.
func CipherSuitesTLS() []uint16 {
	return cipherSuitesTLS()
}

// EllipticCurvesTLS returns the supported elliptic
// curves used by the TLS stack.
func EllipticCurvesTLS() []tls.CurveID {
	return ellipticCurvesTLS()
}
