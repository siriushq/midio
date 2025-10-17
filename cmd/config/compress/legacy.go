package compress

import (
	"strings"

	"github.com/siriushq/midio/cmd/config"
)

// Legacy envs.
const (
	EnvCompress                = "MINIO_COMPRESS"
	EnvCompressMimeTypesLegacy = "MINIO_COMPRESS_MIMETYPES"
)

// SetCompressionConfig - One time migration code needed, for migrating from older config to new for Compression.
func SetCompressionConfig(s config.Config, cfg Config) {
	if !cfg.Enabled {
		// No need to save disabled settings in new config.
		return
	}
	s[config.CompressionSubSys][config.Default] = config.KVS{
		config.KV{
			Key:   config.Enable,
			Value: config.EnableOn,
		},
		config.KV{
			Key:   Extensions,
			Value: strings.Join(cfg.Extensions, config.ValueSeparator),
		},
		config.KV{
			Key:   MimeTypes,
			Value: strings.Join(cfg.MimeTypes, config.ValueSeparator),
		},
	}
}
