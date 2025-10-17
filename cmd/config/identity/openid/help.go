package openid

import "github.com/siriushq/midio/cmd/config"

// Help template for OpenID identity feature.
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         ConfigURL,
			Description: `openid discovery document e.g. "https://accounts.google.com/.well-known/openid-configuration"`,
			Type:        "url",
		},
		config.HelpKV{
			Key:         ClientID,
			Description: `unique public identifier for apps e.g. "292085223830.apps.googleusercontent.com"`,
			Type:        "string",
			Optional:    true,
		},
		config.HelpKV{
			Key:         ClaimName,
			Description: `JWT canned policy claim name, defaults to "policy"`,
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         ClaimPrefix,
			Description: `JWT claim namespace prefix e.g. "customer1/"`,
			Optional:    true,
			Type:        "string",
		},
		config.HelpKV{
			Key:         Scopes,
			Description: `Comma separated list of OpenID scopes for server, defaults to advertised scopes from discovery document e.g. "email,admin"`,
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
