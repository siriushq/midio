package parquet

import "encoding/xml"

// ReaderArgs - represents elements inside <InputSerialization><Parquet/> in request XML.
type ReaderArgs struct {
	unmarshaled bool
}

// IsEmpty - returns whether reader args is empty or not.
func (args *ReaderArgs) IsEmpty() bool {
	return !args.unmarshaled
}

// UnmarshalXML - decodes XML data.
func (args *ReaderArgs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Make subtype to avoid recursive UnmarshalXML().
	type subReaderArgs ReaderArgs
	parsedArgs := subReaderArgs{}
	if err := d.DecodeElement(&parsedArgs, &start); err != nil {
		return err
	}

	args.unmarshaled = true
	return nil
}
