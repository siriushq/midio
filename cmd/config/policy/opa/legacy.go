package opa

import (
	"github.com/siriushq/midio/cmd/config"
)

// Legacy OPA envs
const (
	EnvIamOpaURL       = "MINIO_IAM_OPA_URL"
	EnvIamOpaAuthToken = "MINIO_IAM_OPA_AUTHTOKEN"
)

// SetPolicyOPAConfig - One time migration code needed, for migrating from older config to new for PolicyOPAConfig.
func SetPolicyOPAConfig(s config.Config, opaArgs Args) {
	if opaArgs.URL == nil || opaArgs.URL.String() == "" {
		// Do not enable if opaArgs was empty.
		return
	}
	s[config.PolicyOPASubSys][config.Default] = config.KVS{
		config.KV{
			Key:   URL,
			Value: opaArgs.URL.String(),
		},
		config.KV{
			Key:   AuthToken,
			Value: opaArgs.AuthToken,
		},
	}
}
