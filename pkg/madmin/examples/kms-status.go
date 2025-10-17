//go:build ignore
// +build ignore

package main

import (
	"context"
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

	status, err := madmClnt.GetKeyStatus(context.Background(), "") // empty string refers to the default master key
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Key: %s\n", status.KeyID)
	if status.EncryptionErr == "" {
		log.Println("\t • Encryption ✔")
	} else {
		log.Printf("\t • Encryption failed: %s\n", status.EncryptionErr)
	}
	if status.UpdateErr == "" {
		log.Println("\t • Re-wrap ✔")
	} else {
		log.Printf("\t • Re-wrap failed: %s\n", status.UpdateErr)
	}
	if status.DecryptionErr == "" {
		log.Println("\t • Decryption ✔")
	} else {
		log.Printf("\t •  Decryption failed: %s\n", status.DecryptionErr)
	}
}
