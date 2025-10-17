package storageclass

import (
	"github.com/siriushq/midio/cmd/config"
)

// SetStorageClass - One time migration code needed, for migrating from older config to new for StorageClass.
func SetStorageClass(s config.Config, cfg Config) {
	if len(cfg.Standard.String()) == 0 && len(cfg.RRS.String()) == 0 {
		// Do not enable storage-class if no settings found.
		return
	}
	s[config.StorageClassSubSys][config.Default] = config.KVS{
		config.KV{
			Key:   ClassStandard,
			Value: cfg.Standard.String(),
		},
		config.KV{
			Key:   ClassRRS,
			Value: cfg.RRS.String(),
		},
	}
}
