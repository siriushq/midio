//go:build ignore
// +build ignore

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/minio/minio-go/v7"
	cr "github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	// LDAP integrated Minio endpoint
	stsEndpoint string

	// LDAP credentials
	ldapUsername string
	ldapPassword string
)

func init() {
	flag.StringVar(&stsEndpoint, "sts-ep", "http://localhost:9000", "STS endpoint")
	flag.StringVar(&ldapUsername, "u", "", "AD/LDAP Username")
	flag.StringVar(&ldapPassword, "p", "", "AD/LDAP Password")
}

func main() {
	flag.Parse()
	if ldapUsername == "" || ldapPassword == "" {
		flag.PrintDefaults()
		return
	}

	// The credentials package in minio-go provides an interface to call the
	// LDAP STS API.

	// Initialize LDAP credentials
	li, _ := cr.NewLDAPIdentity(stsEndpoint, ldapUsername, ldapPassword)

	stsEndpointURL, err := url.Parse(stsEndpoint)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	opts := &minio.Options{
		Creds:  li,
		Secure: stsEndpointURL.Scheme == "https",
	}

	fmt.Println(li.Get())
	// Use generated credentials to authenticate with MinIO server
	minioClient, err := minio.New(stsEndpointURL.Host, opts)
	if err != nil {
		log.Fatalln(err)
	}

	// Use minIO Client object normally like the regular client.
	fmt.Println("Calling list objects with temp creds: ")
	objCh := minioClient.ListObjects(context.Background(), ldapUsername, minio.ListObjectsOptions{})
	for obj := range objCh {
		if obj.Err != nil {
			if err != nil {
				log.Fatalln(err)
			}
		}
		fmt.Println(obj)
	}
}
