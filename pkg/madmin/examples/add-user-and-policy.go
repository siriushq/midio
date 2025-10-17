//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"

	"github.com/siriushq/midio/pkg/bucket/policy"
	"github.com/siriushq/midio/pkg/bucket/policy/condition"
	iampolicy "github.com/siriushq/midio/pkg/iam/policy"
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

	if err = madmClnt.AddUser(context.Background(), "newuser", "newstrongpassword"); err != nil {
		log.Fatalln(err)
	}

	// Create policy
	p := iampolicy.Policy{
		Version: iampolicy.DefaultVersion,
		Statements: []iampolicy.Statement{
			iampolicy.NewStatement(
				policy.Allow,
				iampolicy.NewActionSet(iampolicy.GetObjectAction),
				iampolicy.NewResourceSet(iampolicy.NewResource("testbucket/*", "")),
				condition.NewFunctions(),
			)},
	}

	if err = madmClnt.AddCannedPolicy(context.Background(), "get-only", &p); err != nil {
		log.Fatalln(err)
	}

	if err = madmClnt.SetUserPolicy(context.Background(), "newuser", "get-only"); err != nil {
		log.Fatalln(err)
	}
}
