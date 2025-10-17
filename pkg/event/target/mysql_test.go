package target

import (
	"database/sql"
	"testing"
)

// TestPostgreSQLRegistration checks if sql driver
// is registered and fails otherwise.
func TestMySQLRegistration(t *testing.T) {
	var found bool
	for _, drv := range sql.Drivers() {
		if drv == "mysql" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("mysql driver not registered")
	}
}
