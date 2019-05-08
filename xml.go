package sedoc

import (
	"encoding/xml"
	"fmt"
	"sort"
)

// MarshallerXML for marshal []interface{} into XML
func MarshallerXML(arr []interface{}) func(*xml.Encoder, xml.StartElement) error {
	return func(e *xml.Encoder, start xml.StartElement) error {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Space: "", Local: "count"},
			Value: fmt.Sprint(len(arr)),
		})
		if err := e.EncodeToken(start); err != nil {
			return err
		}
		for _, value := range arr {
			if err := e.Encode(value); err != nil {
				return err
			}
		}
		if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
			return err
		}
		// flush to ensure tokens are written
		err := e.Flush()
		return err
	}
}

// StringMap is a map[string]string
type StringMap map[string]string

// MarshalXML marshals StringMap into XML
func (s StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := s[key]
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: key}, Value: value})
	}
	tokens := []xml.Token{start, start.End()}
	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}
	// flush to ensure tokens are written
	err := e.Flush()
	return err
}

// UnmarshalXML marshals StringMap into XML
func (s *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*s = StringMap{}
	for _, attr := range start.Attr {
		(*s)[attr.Name.Local] = attr.Value
	}
	return d.Skip()
}

// InterfaceMap is a map[string]interface{}
type InterfaceMap map[string]interface{}

// MarshalXML marshals InterfaceMap into XML
func (s InterfaceMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		var value string
		if _, ok := s[key].(string); ok {
			value = s[key].(string)
		} else {
			value = fmt.Sprintf("%v", s[key])
			// return errors.New("MarshalXML: invalid value type (field '" + key + "' typeof '" + fmt.Sprintf("%T", value) + "')")
		}
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: key}, Value: value})
	}
	tokens := []xml.Token{start, start.End()}
	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}
	// flush to ensure tokens are written
	err := e.Flush()
	return err
}

// UnmarshalXML marshals InterfaceMap into XML
func (s *InterfaceMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*s = InterfaceMap{}
	for _, attr := range start.Attr {
		(*s)[attr.Name.Local] = attr.Value
	}
	return d.Skip()
}
