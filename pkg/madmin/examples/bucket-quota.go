//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/siriushq/midio/pkg/madmin"
)

func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY and my-bucketname are
	// dummy values, please replace them with original values.

	// API requests are secure (HTTPS) if secure=true and insecure (HTTP) otherwise.
	// New returns an MinIO Admin client object.
	madmClnt, err := madmin.New("your-minio.example.com:9000", "YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", true)
	if err != nil {
		log.Fatalln(err)
	}
	var kiB int64 = 1 << 10
	ctx := context.Background()
	// set bucket quota config
	if err := madmClnt.SetBucketQuota(ctx, "bucket-name", 64*kiB, HardQuota); err != nil {
		log.Fatalln(err)
	}
	// gets bucket quota config
	quotaCfg, err := madmClnt.GetBucketQuota(ctx, "bucket-name")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(quotaCfg)
	// remove bucket quota config
	if err := madmClnt.RemoveBucketQuota(ctx, "bucket-name"); err != nil {
		log.Fatalln(err)
	}
}
