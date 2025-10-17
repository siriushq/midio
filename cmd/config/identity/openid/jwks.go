package openid

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
)

// JWKS - https://tools.ietf.org/html/rfc7517
type JWKS struct {
	Keys []*JWKS `json:"keys,omitempty"`

	Kty string `json:"kty"`
	Use string `json:"use,omitempty"`
	Kid string `json:"kid,omitempty"`
	Alg string `json:"alg,omitempty"`

	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Y   string `json:"y,omitempty"`
	D   string `json:"d,omitempty"`
	N   string `json:"n,omitempty"`
	E   string `json:"e,omitempty"`
	K   string `json:"k,omitempty"`
}

var (
	errMalformedJWKRSAKey = errors.New("malformed JWK RSA key")
	errMalformedJWKECKey  = errors.New("malformed JWK EC key")
)

// DecodePublicKey - decodes JSON Web Key (JWK) as public key
func (key *JWKS) DecodePublicKey() (crypto.PublicKey, error) {
	switch key.Kty {
	case "RSA":
		if key.N == "" || key.E == "" {
			return nil, errMalformedJWKRSAKey
		}

		// decode exponent
		ebuf, err := base64.RawURLEncoding.DecodeString(key.E)
		if err != nil {
			return nil, errMalformedJWKRSAKey
		}

		nbuf, err := base64.RawURLEncoding.DecodeString(key.N)
		if err != nil {
			return nil, errMalformedJWKRSAKey
		}

		var n, e big.Int
		n.SetBytes(nbuf)
		e.SetBytes(ebuf)

		return &rsa.PublicKey{
			E: int(e.Int64()),
			N: &n,
		}, nil
	case "EC":
		if key.Crv == "" || key.X == "" || key.Y == "" {
			return nil, errMalformedJWKECKey
		}

		var curve elliptic.Curve
		switch key.Crv {
		case "P-224":
			curve = elliptic.P224()
		case "P-256":
			curve = elliptic.P256()
		case "P-384":
			curve = elliptic.P384()
		case "P-521":
			curve = elliptic.P521()
		default:
			return nil, fmt.Errorf("Unknown curve type: %s", key.Crv)
		}

		xbuf, err := base64.RawURLEncoding.DecodeString(key.X)
		if err != nil {
			return nil, errMalformedJWKECKey
		}

		ybuf, err := base64.RawURLEncoding.DecodeString(key.Y)
		if err != nil {
			return nil, errMalformedJWKECKey
		}

		var x, y big.Int
		x.SetBytes(xbuf)
		y.SetBytes(ybuf)

		return &ecdsa.PublicKey{
			Curve: curve,
			X:     &x,
			Y:     &y,
		}, nil
	default:
		return nil, fmt.Errorf("Unknown JWK key type %s", key.Kty)
	}
}
