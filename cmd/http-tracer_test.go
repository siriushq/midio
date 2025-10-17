package cmd

import (
	"testing"
)

// Test redactLDAPPwd()
func TestRedactLDAPPwd(t *testing.T) {
	testCases := []struct {
		query         string
		expectedQuery string
	}{
		{"", ""},
		{"?Action=AssumeRoleWithLDAPIdentity&LDAPUsername=myusername&LDAPPassword=can+youreadthis%3F&Version=2011-06-15",
			"?Action=AssumeRoleWithLDAPIdentity&LDAPUsername=myusername&LDAPPassword=*REDACTED*&Version=2011-06-15",
		},
		{"LDAPPassword=can+youreadthis%3F&Version=2011-06-15&?Action=AssumeRoleWithLDAPIdentity&LDAPUsername=myusername",
			"LDAPPassword=*REDACTED*&Version=2011-06-15&?Action=AssumeRoleWithLDAPIdentity&LDAPUsername=myusername",
		},
		{"?Action=AssumeRoleWithLDAPIdentity&LDAPUsername=myusername&Version=2011-06-15&LDAPPassword=can+youreadthis%3F",
			"?Action=AssumeRoleWithLDAPIdentity&LDAPUsername=myusername&Version=2011-06-15&LDAPPassword=*REDACTED*",
		},
		{
			"?x=y&a=b",
			"?x=y&a=b",
		},
	}
	for i, test := range testCases {
		gotQuery := redactLDAPPwd(test.query)
		if gotQuery != test.expectedQuery {
			t.Fatalf("test %d: expected %s got %s", i+1, test.expectedQuery, gotQuery)
		}
	}
}
