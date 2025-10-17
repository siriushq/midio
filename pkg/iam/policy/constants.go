package iampolicy

import (
	"github.com/siriushq/midio/pkg/bucket/policy"
	"github.com/siriushq/midio/pkg/bucket/policy/condition"
)

// Policy claim constants
const (
	PolicyName        = "policy"
	SessionPolicyName = "sessionPolicy"
)

// ReadWrite - provides full access to all buckets and all objects
var ReadWrite = Policy{
	Version: DefaultVersion,
	Statements: []Statement{
		{
			SID:       policy.ID(""),
			Effect:    policy.Allow,
			Actions:   NewActionSet(AllActions),
			Resources: NewResourceSet(NewResource("*", "")),
		},
	},
}

// ReadOnly - read only.
var ReadOnly = Policy{
	Version: DefaultVersion,
	Statements: []Statement{
		{
			SID:       policy.ID(""),
			Effect:    policy.Allow,
			Actions:   NewActionSet(GetBucketLocationAction, GetObjectAction),
			Resources: NewResourceSet(NewResource("*", "")),
		},
	},
}

// WriteOnly - provides write access.
var WriteOnly = Policy{
	Version: DefaultVersion,
	Statements: []Statement{
		{
			SID:       policy.ID(""),
			Effect:    policy.Allow,
			Actions:   NewActionSet(PutObjectAction),
			Resources: NewResourceSet(NewResource("*", "")),
		},
	},
}

// AdminDiagnostics - provides admin diagnostics access.
var AdminDiagnostics = Policy{
	Version: DefaultVersion,
	Statements: []Statement{
		{
			SID:    policy.ID(""),
			Effect: policy.Allow,
			Actions: NewActionSet(ProfilingAdminAction,
				TraceAdminAction, ConsoleLogAdminAction,
				ServerInfoAdminAction, TopLocksAdminAction,
				HealthInfoAdminAction, BandwidthMonitorAction,
				PrometheusAdminAction,
			),
			Resources: NewResourceSet(NewResource("*", "")),
		},
	},
}

// Admin - provides admin all-access canned policy
var Admin = Policy{
	Version: DefaultVersion,
	Statements: []Statement{
		{
			SID:        policy.ID(""),
			Effect:     policy.Allow,
			Actions:    NewActionSet(AllAdminActions),
			Resources:  NewResourceSet(),
			Conditions: condition.NewFunctions(),
		},
		{
			SID:        policy.ID(""),
			Effect:     policy.Allow,
			Actions:    NewActionSet(AllActions),
			Resources:  NewResourceSet(NewResource("*", "")),
			Conditions: condition.NewFunctions(),
		},
	},
}
