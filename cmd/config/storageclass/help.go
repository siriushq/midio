package storageclass

import "github.com/siriushq/midio/cmd/config"

// Help template for storageclass feature.
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         ClassStandard,
			Description: `set the parity count for default standard storage class e.g. "EC:4"`,
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         ClassRRS,
			Description: `set the parity count for reduced redundancy storage class e.g. "EC:2"`,
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         ClassDMA,
			Description: `enable O_DIRECT for both read and write, defaults to "write" e.g. "read+write"`,
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
