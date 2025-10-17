package cmd

// SetupType - enum for setup type.
type SetupType int

const (
	// UnknownSetupType - starts with unknown setup type.
	UnknownSetupType SetupType = iota

	// FSSetupType - FS setup type enum.
	FSSetupType

	// ErasureSetupType - Erasure setup type enum.
	ErasureSetupType

	// DistErasureSetupType - Distributed Erasure setup type enum.
	DistErasureSetupType

	// GatewaySetupType - gateway setup type enum.
	GatewaySetupType
)

func (setupType SetupType) String() string {
	switch setupType {
	case FSSetupType:
		return globalMinioModeFS
	case ErasureSetupType:
		return globalMinioModeErasure
	case DistErasureSetupType:
		return globalMinioModeDistErasure
	case GatewaySetupType:
		return globalMinioModeGatewayPrefix
	}

	return "unknown"
}
