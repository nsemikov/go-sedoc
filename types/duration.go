package types

import (
	"encoding/xml"
	"strings"
	"time"
)

// Duration type
type Duration time.Duration

// String method
func (d Duration) String() string {
	return time.Duration(d).String()
}

// -----------------------------------------------------------------------------

// UnmarshalJSON method
func (d *Duration) UnmarshalJSON(b []byte) error {
	t, err := time.ParseDuration(strings.Trim(string(b), `"`))
	*d = Duration(t)
	return err
}

// MarshalJSON method
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

// -----------------------------------------------------------------------------

// UnmarshalXML method
func (d *Duration) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var s string
	err := dec.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	var t time.Duration
	t, err = time.ParseDuration(s)
	if err == nil {
		*d = Duration(t)
	}
	return err
}

// UnmarshalXMLAttr method
func (d *Duration) UnmarshalXMLAttr(attr xml.Attr) error {
	t, err := time.ParseDuration(attr.Value)
	if err == nil {
		*d = Duration(t)
	}
	return err
}

// MarshalXML method
func (d Duration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(d.String(), start)
}

// MarshalXMLAttr method
func (d Duration) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: d.String(),
	}, nil
}

// -----------------------------------------------------------------------------

// UnmarshalYAML method
func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	t, err := time.ParseDuration(s)
	if err == nil {
		*d = Duration(t)
	}
	return err
}

// MarshalYAML method
func (d Duration) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}
