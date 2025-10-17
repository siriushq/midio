package compress

import "github.com/siriushq/midio/cmd/config"

// Help template for compress feature.
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         Extensions,
			Description: `comma separated file extensions e.g. ".txt,.log,.csv"`,
			Optional:    true,
			Type:        "csv",
		},
		config.HelpKV{
			Key:         MimeTypes,
			Description: `comma separated wildcard mime-types e.g. "text/*,application/json,application/xml"`,
			Optional:    true,
			Type:        "csv",
		},
		config.HelpKV{
			Key:         config.Comment,
			Description: config.DefaultComment,
			Optional:    true,
			Type:        "sentence",
		},
	}
)
