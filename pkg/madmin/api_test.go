// Package madmin_test
package madmin_test

import (
	"testing"

	"github.com/siriushq/midio/pkg/madmin"
)

func TestMinioAdminClient(t *testing.T) {
	_, err := madmin.New("localhost:9000", "food", "food123", true)
	if err != nil {
		t.Fatal(err)
	}
}
