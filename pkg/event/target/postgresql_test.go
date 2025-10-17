package target

import (
	"database/sql"
	"testing"
)

// TestPostgreSQLRegistration checks if postgres driver
// is registered and fails otherwise.
func TestPostgreSQLRegistration(t *testing.T) {
	var found bool
	for _, drv := range sql.Drivers() {
		if drv == "postgres" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("postgres driver not registered")
	}
}
