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
	madminClient, err := madmin.New("your-minio.example.com:9000", "YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", true)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	reportCh := madminClient.GetBucketBandwidth(ctx)

	for i := 0; i < 10; i++ {
		report := <-reportCh
		fmt.Printf("Report: %+v\n", report)
	}
	reportCh = madminClient.GetBucketBandwidth(ctx, "sourceBucket", "sourceBucket2")
	for i := 0; i < 10; i++ {
		report := <-reportCh
		fmt.Printf("Report: %+v\n", report)
	}
}
