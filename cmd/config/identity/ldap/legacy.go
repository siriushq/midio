package ldap

import "github.com/siriushq/midio/cmd/config"

// SetIdentityLDAP - One time migration code needed, for migrating from older config to new for LDAPConfig.
func SetIdentityLDAP(s config.Config, ldapArgs Config) {
	if !ldapArgs.Enabled {
		// ldap not enabled no need to preserve it in new settings.
		return
	}
	s[config.IdentityLDAPSubSys][config.Default] = config.KVS{
		config.KV{
			Key:   ServerAddr,
			Value: ldapArgs.ServerAddr,
		},
		config.KV{
			Key:   STSExpiry,
			Value: ldapArgs.STSExpiryDuration,
		},
		config.KV{
			Key:   UsernameFormat,
			Value: ldapArgs.UsernameFormat,
		},
		config.KV{
			Key:   GroupSearchFilter,
			Value: ldapArgs.GroupSearchFilter,
		},
		config.KV{
			Key:   GroupSearchBaseDN,
			Value: ldapArgs.GroupSearchBaseDistName,
		},
	}
}
