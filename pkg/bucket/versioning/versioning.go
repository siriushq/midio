package versioning

import (
	"encoding/xml"
	"io"
)

// State - enabled/disabled/suspended states
// for multifactor and status of versioning.
type State string

// Various supported states
const (
	Enabled State = "Enabled"
	// Disabled  State = "Disabled" only used by MFA Delete not supported yet.
	Suspended State = "Suspended"
)

// Versioning - Configuration for bucket versioning.
type Versioning struct {
	XMLNS   string   `xml:"xmlns,attr,omitempty"`
	XMLName xml.Name `xml:"VersioningConfiguration"`
	// MFADelete State    `xml:"MFADelete,omitempty"` // not supported yet.
	Status State `xml:"Status,omitempty"`
}

// Validate - validates the versioning configuration
func (v Versioning) Validate() error {
	// Not supported yet
	// switch v.MFADelete {
	// case Enabled, Disabled:
	// default:
	// 	return Errorf("unsupported MFADelete state %s", v.MFADelete)
	// }
	switch v.Status {
	case Enabled, Suspended:
	default:
		return Errorf("unsupported Versioning status %s", v.Status)
	}
	return nil
}

// Enabled - returns true if versioning is enabled
func (v Versioning) Enabled() bool {
	return v.Status == Enabled
}

// Suspended - returns true if versioning is suspended
func (v Versioning) Suspended() bool {
	return v.Status == Suspended
}

// ParseConfig - parses data in given reader to VersioningConfiguration.
func ParseConfig(reader io.Reader) (*Versioning, error) {
	var v Versioning
	if err := xml.NewDecoder(reader).Decode(&v); err != nil {
		return nil, err
	}
	if err := v.Validate(); err != nil {
		return nil, err
	}
	return &v, nil
}
