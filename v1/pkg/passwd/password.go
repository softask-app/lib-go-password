package passwd

import "encoding/xml"

const (
	value = "***"
)

// Password defines a password string that should not be logged or printed
// normally.
//
// Password implements fmt.Stringer, json.Marshaler, xml.Marshaler,
// xml.MarshalerAttr, and encoding.TextMarshaler.  For each of these
// implementations the Password is encoded as "***".
//
// This means you have to downcast to string if you actually want to print or
// encode the value.
type Password string

func (Password) String() string {
	return value
}

func (p Password) MarshalJSON() ([]byte, error) {
	return []byte(`"` + value + `"`), nil
}

func (Password) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(value, start)
}

func (Password) MarshalText() (text []byte, err error) {
	return []byte(value), nil
}

func (Password) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: value}, nil
}
