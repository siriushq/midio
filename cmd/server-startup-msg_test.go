package cmd

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/siriushq/midio/pkg/madmin"
)

// Tests if we generate storage info.
func TestStorageInfoMsg(t *testing.T) {
	infoStorage := StorageInfo{}
	infoStorage.Disks = []madmin.Disk{
		{Endpoint: "http://127.0.0.1:9000/data/1/", State: madmin.DriveStateOk},
		{Endpoint: "http://127.0.0.1:9000/data/2/", State: madmin.DriveStateOk},
		{Endpoint: "http://127.0.0.1:9000/data/3/", State: madmin.DriveStateOk},
		{Endpoint: "http://127.0.0.1:9000/data/4/", State: madmin.DriveStateOk},
		{Endpoint: "http://127.0.0.1:9001/data/1/", State: madmin.DriveStateOk},
		{Endpoint: "http://127.0.0.1:9001/data/2/", State: madmin.DriveStateOk},
		{Endpoint: "http://127.0.0.1:9001/data/3/", State: madmin.DriveStateOk},
		{Endpoint: "http://127.0.0.1:9001/data/4/", State: madmin.DriveStateOffline},
	}
	infoStorage.Backend.Type = madmin.Erasure

	if msg := getStorageInfoMsg(infoStorage); !strings.Contains(msg, "7 Online, 1 Offline") {
		t.Fatal("Unexpected storage info message, found:", msg)
	}
}

// Tests stripping standard ports from apiEndpoints.
func TestStripStandardPorts(t *testing.T) {
	apiEndpoints := []string{"http://127.0.0.1:9000", "http://127.0.0.2:80", "https://127.0.0.3:443"}
	expectedAPIEndpoints := []string{"http://127.0.0.1:9000", "http://127.0.0.2", "https://127.0.0.3"}
	newAPIEndpoints := stripStandardPorts(apiEndpoints)

	if !reflect.DeepEqual(expectedAPIEndpoints, newAPIEndpoints) {
		t.Fatalf("Expected %#v, got %#v", expectedAPIEndpoints, newAPIEndpoints)
	}

	apiEndpoints = []string{"http://%%%%%:9000"}
	newAPIEndpoints = stripStandardPorts(apiEndpoints)
	if !reflect.DeepEqual([]string{""}, newAPIEndpoints) {
		t.Fatalf("Expected %#v, got %#v", apiEndpoints, newAPIEndpoints)
	}

	apiEndpoints = []string{"http://127.0.0.1:443", "https://127.0.0.1:80"}
	newAPIEndpoints = stripStandardPorts(apiEndpoints)
	if !reflect.DeepEqual(apiEndpoints, newAPIEndpoints) {
		t.Fatalf("Expected %#v, got %#v", apiEndpoints, newAPIEndpoints)
	}
}

// Test printing server common message.
func TestPrintServerCommonMessage(t *testing.T) {
	obj, fsDir, err := prepareFS()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(fsDir)
	if err = newTestConfig(globalMinioDefaultRegion, obj); err != nil {
		t.Fatal(err)
	}

	apiEndpoints := []string{"http://127.0.0.1:9000"}
	printServerCommonMsg(apiEndpoints)
}

// Tests print cli access message.
func TestPrintCLIAccessMsg(t *testing.T) {
	obj, fsDir, err := prepareFS()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(fsDir)
	if err = newTestConfig(globalMinioDefaultRegion, obj); err != nil {
		t.Fatal(err)
	}

	apiEndpoints := []string{"http://127.0.0.1:9000"}
	printCLIAccessMsg(apiEndpoints[0], "myminio")
}

// Test print startup message.
func TestPrintStartupMessage(t *testing.T) {
	obj, fsDir, err := prepareFS()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(fsDir)
	if err = newTestConfig(globalMinioDefaultRegion, obj); err != nil {
		t.Fatal(err)
	}

	apiEndpoints := []string{"http://127.0.0.1:9000"}
	printStartupMessage(apiEndpoints, nil)
}
