//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
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

	// Create a new service account
	creds, err := madmClnt.AddServiceAccount(context.Background(), madmin.AddServiceAccountReq{Policy: &p})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(creds)

	// List all services accounts
	list, err := madmClnt.ListServiceAccounts(context.Background(), "")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(list)

	// Delete a service account
	err = madmClnt.DeleteServiceAccount(context.Background(), list.Accounts[0])
	if err != nil {
		log.Fatalln(err)
	}
}
