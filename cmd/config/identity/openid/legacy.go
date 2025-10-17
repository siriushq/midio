package openid

import "github.com/siriushq/midio/cmd/config"

// Legacy envs
const (
	EnvIamJwksURL = "MINIO_IAM_JWKS_URL"
)

// SetIdentityOpenID - One time migration code needed, for migrating from older config to new for OpenIDConfig.
func SetIdentityOpenID(s config.Config, cfg Config) {
	if cfg.JWKS.URL == nil || cfg.JWKS.URL.String() == "" {
		// No need to save not-enabled settings in new config.
		return
	}
	s[config.IdentityOpenIDSubSys][config.Default] = config.KVS{
		config.KV{
			Key:   JwksURL,
			Value: cfg.JWKS.URL.String(),
		},
		config.KV{
			Key:   ConfigURL,
			Value: "",
		},
		config.KV{
			Key:   ClaimPrefix,
			Value: "",
		},
	}
}
