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
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY are
	// dummy values, please replace them with original values.

	// API requests are secure (HTTPS) if secure=true and insecure (HTTP) otherwise.
	// New returns an MinIO Admin client object.
	madmClnt, err := madmin.New("your-minio.example.com:9000", "YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", true)
	if err != nil {
		log.Fatalln(err)
	}

	// Start listening on all http trace activity from all servers in the minio cluster.
	traceCh := madmClnt.ServiceTrace(context.Background(), madmin.ServiceTraceOpts{
		S3:        true,
		Internal:  true,
		Storage:   true,
		OS:        true,
		Threshold: 0,
	})
	for traceInfo := range traceCh {
		if traceInfo.Err != nil {
			fmt.Println(traceInfo.Err)
		}
		fmt.Println(traceInfo)
	}
}
