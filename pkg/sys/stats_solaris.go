package sys

import "errors"

// GetStats - stub implementation for Solaris, this will not give us
// complete functionality but will enable fs setups on Solaris.
func GetStats() (stats Stats, err error) {
	return Stats{}, errors.New("Not implemented")
}
