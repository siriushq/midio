package net

import (
	"errors"
	"strconv"
)

// Port - network port
type Port uint16

// String - returns string representation of port.
func (p Port) String() string {
	return strconv.Itoa(int(p))
}

// ParsePort - parses string into Port
func ParsePort(s string) (p Port, err error) {
	if s == "https" {
		return Port(443), nil
	} else if s == "http" {
		return Port(80), nil
	}

	var i int
	if i, err = strconv.Atoi(s); err != nil {
		return p, errors.New("invalid port number")
	}

	if i < 0 || i > 65535 {
		return p, errors.New("port must be between 0 to 65535")
	}

	return Port(i), nil
}
