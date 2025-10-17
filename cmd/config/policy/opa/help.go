package opa

import "github.com/siriushq/midio/cmd/config"

// Help template for OPA policy feature.
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         URL,
			Description: `[DEPRECATED] OPA HTTP(s) endpoint e.g. "http://localhost:8181/v1/data/httpapi/authz/allow"`,
			Type:        "url",
		},
		config.HelpKV{
			Key:         AuthToken,
			Description: "[DEPRECATED] authorization token for OPA endpoint",
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         config.Comment,
			Description: config.DefaultComment,
			Optional:    true,
			Type:        "sentence",
		},
	}
)
