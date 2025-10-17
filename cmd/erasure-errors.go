package cmd

import "errors"

// errErasureReadQuorum - did not meet read quorum.
var errErasureReadQuorum = errors.New("Read failed. Insufficient number of disks online")

// errErasureWriteQuorum - did not meet write quorum.
var errErasureWriteQuorum = errors.New("Write failed. Insufficient number of disks online")

// errNoHealRequired - returned when healing is attempted on a previously healed disks.
var errNoHealRequired = errors.New("No healing is required")
