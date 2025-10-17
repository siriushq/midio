package cache

import (
	"fmt"
	"strings"

	"github.com/siriushq/midio/cmd/config"
)

const (
	cacheDelimiterLegacy = ";"
)

// SetCacheConfig - One time migration code needed, for migrating from older config to new for Cache.
func SetCacheConfig(s config.Config, cfg Config) {
	if len(cfg.Drives) == 0 {
		// Do not save cache if no settings available.
		return
	}
	s[config.CacheSubSys][config.Default] = config.KVS{
		config.KV{
			Key:   Drives,
			Value: strings.Join(cfg.Drives, cacheDelimiter),
		},
		config.KV{
			Key:   Exclude,
			Value: strings.Join(cfg.Exclude, cacheDelimiter),
		},
		config.KV{
			Key:   Expiry,
			Value: fmt.Sprintf("%d", cfg.Expiry),
		},
		config.KV{
			Key:   Quota,
			Value: fmt.Sprintf("%d", cfg.MaxUse),
		},
	}
}
