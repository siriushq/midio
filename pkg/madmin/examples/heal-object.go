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

	// Heal object mybucket/myobject - dry run.
	isDryRun := true
	_, err = madmClnt.HealObject(context.Background(), "mybucket", "myobject", isDryRun)
	if err != nil {
		log.Fatalln(err)
	}

	// Heal object mybucket/myobject - this time for real.
	isDryRun = false
	healResult, err := madmClnt.HealObject(context.Background(), "mybucket", "myobject", isDryRun)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("heal result: %#v\n", healResult)
}
