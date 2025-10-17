package logger

import (
	"github.com/siriushq/midio/cmd/config"
)

// Help template for logger http and audit
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         Endpoint,
			Description: `HTTP(s) endpoint e.g. "http://localhost:8080/minio/logs/server"`,
			Type:        "url",
		},
		config.HelpKV{
			Key:         AuthToken,
			Description: `opaque string or JWT authorization token`,
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

	HelpAudit = config.HelpKVS{
		config.HelpKV{
			Key:         Endpoint,
			Description: `HTTP(s) endpoint e.g. "http://localhost:8080/minio/logs/audit"`,
			Type:        "url",
		},
		config.HelpKV{
			Key:         AuthToken,
			Description: `opaque string or JWT authorization token`,
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         config.Comment,
			Description: config.DefaultComment,
			Optional:    true,
			Type:        "sentence",
		},
		config.HelpKV{
			Key:         ClientCert,
			Description: "mTLS certificate for Audit Webhook authentication",
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         ClientKey,
			Description: "mTLS certificate key for Audit Webhook authentication",
			Optional:    true,
			Type:        "string",
		},
	}
)
