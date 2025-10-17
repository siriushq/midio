//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"

	"github.com/siriushq/midio/pkg/madmin"
)

func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY are
	// dummy values, please replace them with original values.

	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY are
	// dummy values, please replace them with original values.

	// API requests are secure (HTTPS) if secure=true and insecure (HTTP) otherwise.
	// New returns an MinIO Admin client object.
	madmClnt, err := madmin.New("your-minio.example.com:9000", "YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", true)
	if err != nil {
		log.Fatalln(err)
	}

	// Attempt healing format in dry-run mode.
	isDryRun := true
	err = madmClnt.HealFormat(context.Background(), isDryRun)
	if err != nil {
		log.Fatalln(err)
	}

	// Perform actual healing of format.
	isDryRun = false
	err = madmClnt.HealFormat(context.Background(), isDryRun)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("successfully healed storage format on available disks.")
}
