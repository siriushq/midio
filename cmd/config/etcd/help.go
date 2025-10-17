package etcd

import "github.com/siriushq/midio/cmd/config"

// etcd config documented in default config
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         Endpoints,
			Description: `comma separated list of etcd endpoints e.g. "http://localhost:2379"`,
			Type:        "csv",
		},
		config.HelpKV{
			Key:         PathPrefix,
			Description: `namespace prefix to isolate tenants e.g. "customer1/"`,
			Optional:    true,
			Type:        "path",
		},
		config.HelpKV{
			Key:         CoreDNSPath,
			Description: `shared bucket DNS records, default is "/skydns"`,
			Optional:    true,
			Type:        "path",
		},
		config.HelpKV{
			Key:         ClientCert,
			Description: `client cert for mTLS authentication`,
			Optional:    true,
			Type:        "path",
		},
		config.HelpKV{
			Key:         ClientCertKey,
			Description: `client cert key for mTLS authentication`,
			Optional:    true,
			Type:        "path",
		},
		config.HelpKV{
			Key:         config.Comment,
			Description: config.DefaultComment,
			Optional:    true,
			Type:        "sentence",
		},
	}
)
