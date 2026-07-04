// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package cmd

import (
	"crypto/x509"
	"fmt"
	"net"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/siriushq/midio/cmd/config"
	"github.com/siriushq/midio/cmd/logger"
	color "github.com/siriushq/midio/pkg/color"
	"github.com/siriushq/midio/pkg/madmin"
	xnet "github.com/siriushq/midio/pkg/net"
)

// Documentation links, these are part of message printing code.
const (
	mcQuickStartGuide     = "https://docs.min.io/docs/minio-client-quickstart-guide"
	goQuickStartGuide     = "https://docs.min.io/docs/golang-client-quickstart-guide"
	jsQuickStartGuide     = "https://docs.min.io/docs/javascript-client-quickstart-guide"
	javaQuickStartGuide   = "https://docs.min.io/docs/java-client-quickstart-guide"
	pyQuickStartGuide     = "https://docs.min.io/docs/python-client-quickstart-guide"
	dotnetQuickStartGuide = "https://docs.min.io/docs/dotnet-client-quickstart-guide"
)

// generates format string depending on the string length and padding.
func getFormatStr(strLen int, padding int) string {
	formatStr := fmt.Sprintf("%ds", strLen+padding)
	return "%" + formatStr
}

func mustGetStorageInfo(objAPI ObjectLayer) StorageInfo {
	storageInfo, _ := objAPI.StorageInfo(GlobalContext)
	return storageInfo
}

// Prints the formatted startup message.
func printStartupMessage(apiEndpoints []string, err error) {
	if err != nil {
		logStartupMessage(color.RedBold("Server startup failed with '%v'", err))
		logStartupMessage(color.RedBold("Not all features may be available on this server"))
		logStartupMessage(color.RedBold("Please use 'mc admin' commands to further investigate this issue"))
	}

	strippedAPIEndpoints := stripStandardPorts(apiEndpoints)
	// If cache layer is enabled, print cache capacity.
	cachedObjAPI := newCachedObjectLayerFn()
	if cachedObjAPI != nil {
		printCacheStorageInfo(cachedObjAPI.StorageInfo(GlobalContext))
	}

	// Object layer is initialized then print StorageInfo.
	objAPI := newObjectLayerFn()
	if objAPI != nil {
		printStorageInfo(mustGetStorageInfo(objAPI))
	}

	// Prints credential, region and browser access.
	printServerCommonMsg(strippedAPIEndpoints)

	// SSL is configured reads certification chain, prints
	// authority and expiry.
	if color.IsTerminal() && !globalCLIContext.Anonymous {
		if globalIsTLS {
			printCertificateMsg(globalPublicCerts)
		}
	}
}

// Returns true if input is not IPv4, false if it is.
func isNotIPv4(host string) bool {
	h, _, err := net.SplitHostPort(host)
	if err != nil {
		h = host
	}
	ip := net.ParseIP(h)
	ok := ip.To4() != nil // This is always true of IP is IPv4

	// Returns true if input is not IPv4.
	return !ok
}

// strip api endpoints list with standard ports such as
// port "80" and "443" before displaying on the startup
// banner.  Returns a new list of API endpoints.
func stripStandardPorts(apiEndpoints []string) (newAPIEndpoints []string) {
	newAPIEndpoints = make([]string, len(apiEndpoints))
	// Check all API endpoints for standard ports and strip them.
	for i, apiEndpoint := range apiEndpoints {
		u, err := xnet.ParseHTTPURL(apiEndpoint)
		if err != nil {
			continue
		}
		if globalMinioHost == "" && isNotIPv4(u.Host) {
			// Skip all non-IPv4 endpoints when we bind to all interfaces.
			continue
		}
		newAPIEndpoints[i] = u.String()
	}
	return newAPIEndpoints
}

// Prints common server startup message. Prints credential, region and browser access.
func printServerCommonMsg(apiEndpoints []string) {
	// Get saved credentials.
	cred := globalActiveCred

	// Get saved region.
	region := globalServerRegion

	var apiEndpointStrBuilder strings.Builder
	apiEndpointStrBuilder.WriteString("\n")
	for i := 0; i < len(apiEndpoints); i++ {
		if strings.TrimSpace(apiEndpoints[i]) == "" {
			continue
		}
		apiEndpointStrBuilder.WriteString(apiEndpoints[i] + "\n")
	}
	apiEndpointStr := apiEndpointStrBuilder.String()

	// Colorize the message and print.
	logStartupMessage(color.Blue("Endpoint: ") + color.Bold(apiEndpointStr))
	if color.IsTerminal() && !globalCLIContext.Anonymous {
		logStartupMessage(color.Blue("Root Access Key: ") + color.Bold(cred.AccessKey))
		logStartupMessage(color.Blue("Root Secret Key: ") + color.Bold(cred.SecretKey))
		if region != "" {
			logStartupMessage(color.Blue("Region: ") + color.Bold(fmt.Sprintf(getFormatStr(len(region), 2), region)))
		}
	}
	printEventNotifiers()
}

// Prints bucket notification configurations.
func printEventNotifiers() {
	if globalNotificationSys == nil {
		return
	}

	arns := globalNotificationSys.GetARNList(true)
	if len(arns) == 0 {
		return
	}

	arnMsg := color.Blue("SQS ARNs: ")
	for _, arn := range arns {
		arnMsg += color.Bold(fmt.Sprintf("%s ", arn))
	}

	logStartupMessage(arnMsg)
}

// Get formatted disk/storage info message.
func getStorageInfoMsg(storageInfo StorageInfo) string {
	var msg string
	var mcMessage string
	onlineDisks, offlineDisks := getOnlineOfflineDisksStats(storageInfo.Disks)
	if storageInfo.Backend.Type == madmin.Erasure {
		if offlineDisks.Sum() > 0 {
			mcMessage = "Use `mc admin info` to look for latest server/disk info\n"
		}

		diskInfo := fmt.Sprintf(" %d Online, %d Offline. ", onlineDisks.Sum(), offlineDisks.Sum())
		msg += color.Blue("Status:") + fmt.Sprintf(getFormatStr(len(diskInfo), 8), diskInfo)
		if len(mcMessage) > 0 {
			msg = fmt.Sprintf("%s %s", mcMessage, msg)
		}
	}
	return msg
}

// Prints startup message of storage capacity and erasure information.
func printStorageInfo(storageInfo StorageInfo) {
	if msg := getStorageInfoMsg(storageInfo); msg != "" {
		if globalCLIContext.Quiet {
			logger.Info(msg)
		}
		logStartupMessage(msg)
	}
}

func printCacheStorageInfo(storageInfo CacheStorageInfo) {
	msg := fmt.Sprintf("%s %s Free, %s Total", color.Blue("Cache Capacity:"),
		humanize.IBytes(storageInfo.Free),
		humanize.IBytes(storageInfo.Total))
	logStartupMessage(msg)
}

// Prints the certificate expiry message.
func printCertificateMsg(certs []*x509.Certificate) {
	for _, cert := range certs {
		logStartupMessage(config.CertificateText(cert))
	}
}
