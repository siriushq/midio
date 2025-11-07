// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
package policy

import (
	"unicode/utf8"
)

// ID - policy ID.
type ID string

// IsValid - checks if ID is valid or not.
func (id ID) IsValid() bool {
	return utf8.ValidString(string(id))
}
