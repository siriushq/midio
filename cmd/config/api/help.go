package api

import "github.com/siriushq/midio/cmd/config"

// Help template for storageclass feature.
var (
	Help = config.HelpKVS{
		config.HelpKV{
			Key:         apiRequestsMax,
			Description: `set the maximum number of concurrent requests, e.g. "1600"`,
			Optional:    true,
			Type:        "number",
		},
		config.HelpKV{
			Key:         apiRequestsDeadline,
			Description: `set the deadline for API requests waiting to be processed e.g. "1m"`,
			Optional:    true,
			Type:        "duration",
		},
		config.HelpKV{
			Key:         apiCorsAllowOrigin,
			Description: `set comma separated list of origins allowed for CORS requests e.g. "https://example1.com,https://example2.com"`,
			Optional:    true,
			Type:        "csv",
		},
		config.HelpKV{
			Key:         apiRemoteTransportDeadline,
			Description: `set the deadline for API requests on remote transports while proxying between federated instances e.g. "2h"`,
			Optional:    true,
			Type:        "duration",
		},
		config.HelpKV{
			Key:         apiReplicationWorkers,
			Description: `set the number of replication workers, defaults to 100`,
			Optional:    true,
			Type:        "number",
		},
	}
)
