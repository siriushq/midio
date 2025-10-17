package config

// HelpKV - implements help messages for keys
// with value as description of the keys.
type HelpKV struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Optional    bool   `json:"optional"`

	// Indicates if sub-sys supports multiple targets.
	MultipleTargets bool `json:"multipleTargets"`
}

// HelpKVS - implement order of keys help messages.
type HelpKVS []HelpKV

// Lookup - lookup a key from help kvs.
func (hkvs HelpKVS) Lookup(key string) (HelpKV, bool) {
	for _, hkv := range hkvs {
		if hkv.Key == key {
			return hkv, true
		}
	}
	return HelpKV{}, false
}

// DefaultComment used across all sub-systems.
const DefaultComment = "optionally add a comment to this setting"

// Region and Worm help is documented in default config
var (
	RegionHelp = HelpKVS{
		HelpKV{
			Key:         RegionName,
			Type:        "string",
			Description: `name of the location of the server e.g. "us-west-rack2"`,
			Optional:    true,
		},
		HelpKV{
			Key:         Comment,
			Type:        "sentence",
			Description: DefaultComment,
			Optional:    true,
		},
	}
)
