package cmd

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"github.com/siriushq/midio/cmd/logger"
	"github.com/siriushq/midio/pkg/madmin"
)

// getLocalServerProperty - returns madmin.ServerProperties for only the
// local endpoints from given list of endpoints
func getLocalServerProperty(endpointServerPools EndpointServerPools, r *http.Request) madmin.ServerProperties {
	var localEndpoints Endpoints
	addr := r.Host
	if globalIsDistErasure {
		addr = globalLocalNodeName
	}
	network := make(map[string]string)
	for _, ep := range endpointServerPools {
		for _, endpoint := range ep.Endpoints {
			nodeName := endpoint.Host
			if nodeName == "" {
				nodeName = r.Host
			}
			if endpoint.IsLocal {
				// Only proceed for local endpoints
				network[nodeName] = string(madmin.ItemOnline)
				localEndpoints = append(localEndpoints, endpoint)
				continue
			}
			_, present := network[nodeName]
			if !present {
				if err := isServerResolvable(endpoint, 2*time.Second); err == nil {
					network[nodeName] = string(madmin.ItemOnline)
				} else {
					network[nodeName] = string(madmin.ItemOffline)
					// log once the error
					logger.LogOnceIf(context.Background(), err, nodeName)
				}
			}
		}
	}

	props := madmin.ServerProperties{
		State:    string(madmin.ItemInitializing),
		Endpoint: addr,
		Uptime:   UTCNow().Unix() - globalBootTime.Unix(),
		Version:  Version,
		CommitID: CommitID,
		Network:  network,
	}
	runtime.ReadMemStats(&props.MemStats)

	objLayer := newObjectLayerFn()
	if objLayer != nil && !globalIsGateway {
		// only need Disks information in server mode.
		storageInfo, _ := objLayer.LocalStorageInfo(GlobalContext)
		props.State = string(madmin.ItemOnline)
		props.Disks = storageInfo.Disks
	}

	return props
}
