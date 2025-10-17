package cmd

import (
	"errors"
)

const (
	lockRESTVersion       = "v6" // Add Refresh API
	lockRESTVersionPrefix = SlashSeparator + lockRESTVersion
	lockRESTPrefix        = minioReservedBucketPath + "/lock"
)

const (
	lockRESTMethodHealth      = "/health"
	lockRESTMethodRefresh     = "/refresh"
	lockRESTMethodLock        = "/lock"
	lockRESTMethodRLock       = "/rlock"
	lockRESTMethodUnlock      = "/unlock"
	lockRESTMethodRUnlock     = "/runlock"
	lockRESTMethodForceUnlock = "/force-unlock"

	// lockRESTOwner represents owner UUID
	lockRESTOwner = "owner"

	// Unique ID of lock/unlock request.
	lockRESTUID = "uid"

	// Source contains the line number, function and file name of the code
	// on the client node that requested the lock.
	lockRESTSource = "source"

	// Quroum value to be saved along lock requester info, useful
	// in verifying stale locks
	lockRESTQuorum = "quorum"
)

var (
	errLockConflict       = errors.New("lock conflict")
	errLockNotInitialized = errors.New("lock not initialized")
	errLockNotFound       = errors.New("lock not found")
)
