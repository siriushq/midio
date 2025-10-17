package cmd

import (
	"context"
	"testing"

	"github.com/siriushq/midio/pkg/dsync"
)

// Tests lock rpc client.
func TestLockRESTlient(t *testing.T) {
	endpoint, err := NewEndpoint("http://localhost:9000")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	lkClient := newlockRESTClient(endpoint)
	if !lkClient.IsOnline() {
		t.Fatalf("unexpected error. connection failed")
	}

	// Attempt all calls.
	_, err = lkClient.RLock(context.Background(), dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for Rlock to fail")
	}

	_, err = lkClient.Lock(context.Background(), dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for Lock to fail")
	}

	_, err = lkClient.RUnlock(dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for RUnlock to fail")
	}

	_, err = lkClient.Unlock(dsync.LockArgs{})
	if err == nil {
		t.Fatal("Expected for Unlock to fail")
	}
}
