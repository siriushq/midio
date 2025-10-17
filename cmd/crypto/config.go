package crypto

import (
	"math/rand"
	"strings"

	"github.com/siriushq/midio/cmd/config"
	"github.com/siriushq/midio/pkg/ellipses"
	"github.com/siriushq/midio/pkg/env"
	xnet "github.com/siriushq/midio/pkg/net"
)

const (
	// EnvKMSAutoEncryption is the environment variable used to en/disable
	// SSE-S3 auto-encryption. SSE-S3 auto-encryption, if enabled,
	// requires a valid KMS configuration and turns any non-SSE-C
	// request into an SSE-S3 request.
	// If present EnvAutoEncryption must be either "on" or "off".
	EnvKMSAutoEncryption = "MINIO_KMS_AUTO_ENCRYPTION"
)

// ParseKESEndpoints parses the given endpoint string and
// returns a list of valid endpoint URLs. The order of the
// returned endpoints is randomized.
func ParseKESEndpoints(endpointStr string) ([]string, error) {
	var rawEndpoints []string
	for _, endpoint := range strings.Split(endpointStr, ",") {
		if strings.TrimSpace(endpoint) == "" {
			continue
		}
		if !ellipses.HasEllipses(endpoint) {
			rawEndpoints = append(rawEndpoints, endpoint)
			continue
		}
		pattern, err := ellipses.FindEllipsesPatterns(endpoint)
		if err != nil {
			return nil, Errorf("Invalid KES endpoint %q: %v", endpointStr, err)
		}
		for _, p := range pattern {
			rawEndpoints = append(rawEndpoints, p.Expand()...)
		}
	}
	if len(rawEndpoints) == 0 {
		return nil, Errorf("Invalid KES endpoint %q", endpointStr)
	}

	var (
		randNum   = rand.Intn(len(rawEndpoints))
		endpoints = make([]string, len(rawEndpoints))
	)
	for i, endpoint := range rawEndpoints {
		endpoint, err := xnet.ParseHTTPURL(endpoint)
		if err != nil {
			return nil, Errorf("Invalid KES endpoint %q: %v", endpointStr, err)
		}
		endpoints[(randNum+i)%len(rawEndpoints)] = endpoint.String()
	}
	return endpoints, nil
}

// LookupAutoEncryption returns true if and only if
// the MINIO_KMS_AUTO_ENCRYPTION env. variable is
// set to "on".
func LookupAutoEncryption() bool {
	auto, _ := config.ParseBool(env.Get(EnvKMSAutoEncryption, config.EnableOff))
	return auto
}
