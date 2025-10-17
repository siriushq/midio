//go:build ignore
// +build ignore

package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/siriushq/midio/pkg/madmin"
)

func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY are
	// dummy values, please replace them with original values.

	// API requests are secure (HTTPS) if secure=true and insecure (HTTP) otherwise.
	// New returns an MinIO Admin client object.
	madmClnt, err := madmin.New("your-minio.example.com:9000", "YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", true)
	if err != nil {
		log.Fatalln(err)
	}

	healStatusResult, err := madmClnt.BackgroundHealStatus(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	js, _ := json.MarshalIndent(healStatusResult, "", "  ")

	log.Printf("Heal status result: %s\n", string(js))
}
