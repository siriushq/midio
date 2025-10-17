package crypto

import (
	"github.com/siriushq/midio/pkg/kms"
)

// Context is a list of key-value pairs cryptographically
// associated with a certain object.
type Context = kms.Context

// KMS represents an active and authenticted connection
// to a Key-Management-Service. It supports generating
// data key generation and unsealing of KMS-generated
// data keys.
type KMS = kms.KMS
