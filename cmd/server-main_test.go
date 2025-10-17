package cmd

import (
	"context"
	"reflect"
	"testing"
)

// Tests initializing new object layer.
func TestNewObjectLayer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Tests for FS object layer.
	nDisks := 1
	disks, err := getRandomDisks(nDisks)
	if err != nil {
		t.Fatal("Failed to create disks for the backend")
	}
	defer removeRoots(disks)

	obj, err := newObjectLayer(ctx, mustGetPoolEndpoints(disks...))
	if err != nil {
		t.Fatal("Unexpected object layer initialization error", err)
	}
	_, ok := obj.(*FSObjects)
	if !ok {
		t.Fatal("Unexpected object layer detected", reflect.TypeOf(obj))
	}

	// Tests for Erasure object layer initialization.

	// Create temporary backend for the test server.
	nDisks = 16
	disks, err = getRandomDisks(nDisks)
	if err != nil {
		t.Fatal("Failed to create disks for the backend")
	}
	defer removeRoots(disks)

	obj, err = newObjectLayer(ctx, mustGetPoolEndpoints(disks...))
	if err != nil {
		t.Fatal("Unexpected object layer initialization error", err)
	}

	_, ok = obj.(*erasureServerPools)
	if !ok {
		t.Fatal("Unexpected object layer detected", reflect.TypeOf(obj))
	}
}
