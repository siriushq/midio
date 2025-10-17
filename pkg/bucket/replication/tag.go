package replication

import (
	"encoding/xml"
	"unicode/utf8"
)

// Tag - a tag for a replication configuration Rule filter.
type Tag struct {
	XMLName xml.Name `xml:"Tag" json:"Tag"`
	Key     string   `xml:"Key,omitempty" json:"Key,omitempty"`
	Value   string   `xml:"Value,omitempty" json:"Value,omitempty"`
}

var (
	errInvalidTagKey   = Errorf("The TagKey you have provided is invalid")
	errInvalidTagValue = Errorf("The TagValue you have provided is invalid")
)

func (tag Tag) String() string {
	return tag.Key + "=" + tag.Value
}

// IsEmpty returns whether this tag is empty or not.
func (tag Tag) IsEmpty() bool {
	return tag.Key == ""
}

// Validate checks this tag.
func (tag Tag) Validate() error {
	if len(tag.Key) == 0 || utf8.RuneCountInString(tag.Key) > 128 {
		return errInvalidTagKey
	}

	if utf8.RuneCountInString(tag.Value) > 256 {
		return errInvalidTagValue
	}

	return nil
}
