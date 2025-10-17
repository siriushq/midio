package cmd

import (
	"context"
	"os"
	"testing"

	"github.com/siriushq/midio/cmd/config"
)

func TestServerConfig(t *testing.T) {
	objLayer, fsDir, err := prepareFS()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(fsDir)

	if err = newTestConfig(globalMinioDefaultRegion, objLayer); err != nil {
		t.Fatalf("Init Test config failed")
	}

	if globalServerRegion != globalMinioDefaultRegion {
		t.Errorf("Expecting region `us-east-1` found %s", globalServerRegion)
	}

	// Set new region and verify.
	config.SetRegion(globalServerConfig, "us-west-1")
	region, err := config.LookupRegion(globalServerConfig[config.RegionSubSys][config.Default])
	if err != nil {
		t.Fatal(err)
	}
	if region != "us-west-1" {
		t.Errorf("Expecting region `us-west-1` found %s", globalServerRegion)
	}

	if err := saveServerConfig(context.Background(), objLayer, globalServerConfig); err != nil {
		t.Fatalf("Unable to save updated config file %s", err)
	}

	// Initialize server config.
	if err := loadConfig(objLayer); err != nil {
		t.Fatalf("Unable to initialize from updated config file %s", err)
	}
}
