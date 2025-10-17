package cache

import "github.com/siriushq/midio/cmd/config"

// Help template for caching feature.
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         Drives,
			Description: `comma separated mountpoints e.g. "/optane1,/optane2"`,
			Type:        "csv",
		},
		config.HelpKV{
			Key:         Expiry,
			Description: `cache expiry duration in days e.g. "90"`,
			Optional:    true,
			Type:        "number",
		},
		config.HelpKV{
			Key:         Quota,
			Description: `limit cache drive usage in percentage e.g. "90"`,
			Optional:    true,
			Type:        "number",
		},
		config.HelpKV{
			Key:         Exclude,
			Description: `exclude cache for following patterns e.g. "bucket/*.tmp,*.exe"`,
			Optional:    true,
			Type:        "csv",
		},
		config.HelpKV{
			Key:         After,
			Description: `minimum number of access before caching an object`,
			Optional:    true,
			Type:        "number",
		},
		config.HelpKV{
			Key:         WatermarkLow,
			Description: `% of cache use at which to stop cache eviction`,
			Optional:    true,
			Type:        "number",
		},
		config.HelpKV{
			Key:         WatermarkHigh,
			Description: `% of cache use at which to start cache eviction`,
			Optional:    true,
			Type:        "number",
		},
		config.HelpKV{
			Key:         Range,
			Description: `set to "on" or "off" caching of independent range requests per object, defaults to "on"`,
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         Commit,
			Description: `set to control cache commit behavior, defaults to "writethrough"`,
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
