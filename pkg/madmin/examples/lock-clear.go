//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"
	"time"

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

	// Clear locks held on mybucket/myprefix for longer than 30s.
	olderThan := time.Duration(30 * time.Second)
	locksCleared, err := madmClnt.ClearLocks(context.Background(), "mybucket", "myprefix", olderThan)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(locksCleared)
}
