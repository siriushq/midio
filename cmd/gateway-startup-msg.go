package cmd

import (
	"fmt"
	"strings"

	"github.com/siriushq/midio/pkg/color"
)

// Prints the formatted startup message.
func printGatewayStartupMessage(apiEndPoints []string, backendType string) {
	strippedAPIEndpoints := stripStandardPorts(apiEndPoints)
	// If cache layer is enabled, print cache capacity.
	cacheAPI := newCachedObjectLayerFn()
	if cacheAPI != nil {
		printCacheStorageInfo(cacheAPI.StorageInfo(GlobalContext))
	}
	// Prints credential.
	printGatewayCommonMsg(strippedAPIEndpoints)

	// Prints `mc` cli configuration message chooses
	// first endpoint as default.
	printCLIAccessMsg(strippedAPIEndpoints[0], fmt.Sprintf("my%s", backendType))

	// Prints documentation message.
	printObjectAPIMsg()

	// SSL is configured reads certification chain, prints
	// authority and expiry.
	if color.IsTerminal() && !globalCLIContext.Anonymous {
		if globalIsTLS {
			printCertificateMsg(globalPublicCerts)
		}
	}
}

// Prints common server startup message. Prints credential, region and browser access.
func printGatewayCommonMsg(apiEndpoints []string) {
	// Get saved credentials.
	cred := globalActiveCred

	apiEndpointStr := strings.Join(apiEndpoints, "  ")

	// Colorize the message and print.
	logStartupMessage(color.Blue("Endpoint: ") + color.Bold(fmt.Sprintf("%s ", apiEndpointStr)))
	if color.IsTerminal() && !globalCLIContext.Anonymous {
		logStartupMessage(color.Blue("RootUser: ") + color.Bold(fmt.Sprintf("%s ", cred.AccessKey)))
		logStartupMessage(color.Blue("RootPass: ") + color.Bold(fmt.Sprintf("%s ", cred.SecretKey)))
	}
	printEventNotifiers()

	if globalBrowserEnabled {
		logStartupMessage(color.Blue("\nBrowser Access:"))
		logStartupMessage(fmt.Sprintf(getFormatStr(len(apiEndpointStr), 3), apiEndpointStr))
	}
}
