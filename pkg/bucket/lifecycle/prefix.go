package lifecycle

import (
	"encoding/xml"
)

// Prefix holds the prefix xml tag in <Rule> and <Filter>
type Prefix struct {
	string
	set bool
}

// UnmarshalXML - decodes XML data.
func (p *Prefix) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var s string
	if err = d.DecodeElement(&s, &start); err != nil {
		return err
	}
	*p = Prefix{string: s, set: true}
	return nil
}

// MarshalXML - decodes XML data.
func (p Prefix) MarshalXML(e *xml.Encoder, startElement xml.StartElement) error {
	if !p.set {
		return nil
	}
	return e.EncodeElement(p.string, startElement)
}

// String returns the prefix string
func (p Prefix) String() string {
	return p.string
}
