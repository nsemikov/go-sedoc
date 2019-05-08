package sedoc

import (
	"encoding/json"
	"encoding/xml"

	yaml "gopkg.in/yaml.v2"
)

// Example using for show how to use command
type Example struct {
	XMLName     xml.Name         `json:"-" xml:"example" yaml:"-"`
	Name        string           `json:"name,omitempty" xml:"name,attr,omitempty" yaml:"name,omitempty"`
	Description string           `json:"description,omitempty" xml:"description,attr,omitempty" yaml:"description,omitempty"`
	Request     ExampleRequest   `json:"request,omitempty" xml:"request,omitempty" yaml:"request,omitempty"`
	Responses   ExampleResponses `json:"responses,omitempty" xml:"responses,omitempty" yaml:"responses,omitempty"`
}

// Examples is array of Command
type Examples []Example

// MarshalXML for marshal into XML
func (arr Examples) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	a := []interface{}{}
	for _, item := range arr {
		a = append(a, item)
	}
	return MarshallerXML(a)(e, start)
}

// ExampleRequest using for show how to use command
type ExampleRequest struct {
	XMLName xml.Name `json:"-" xml:"request" yaml:"-"`
	Object  Request  `json:"request" yaml:"request"`
	JSON    string   `json:"json,omitempty" xml:"json,omitempty" yaml:"json,omitempty"`
	XML     string   `json:"xml,omitempty" xml:"xml,omitempty" yaml:"xml,omitempty"`
	YAML    string   `json:"yaml,omitempty" xml:"yaml,omitempty" yaml:"yaml,omitempty"`
}

// XMLString method
func (er ExampleRequest) XMLString() (result string) {
	if b, err := xml.MarshalIndent(er.Object, "", "    "); err == nil {
		result = xml.Header + string(b)
	}
	return
}

// JSONString method
func (er ExampleRequest) JSONString() (result string) {
	if b, err := json.MarshalIndent(er.Object, "", "    "); err == nil {
		result = string(b)
	}
	return
}

// YAMLString method
func (er ExampleRequest) YAMLString() (result string) {
	if b, err := yaml.Marshal(er.Object); err == nil {
		result = string(b)
	}
	return
}

// ExampleResponse using for show how to use command
type ExampleResponse struct {
	XMLName     xml.Name `json:"-" xml:"response" yaml:"-"`
	Name        string   `json:"name,omitempty" xml:"name,attr,omitempty" yaml:"name,omitempty"`
	Description string   `json:"description,omitempty" xml:"description,attr,omitempty" yaml:"description,omitempty"`
	Object      Response `json:"response" yaml:"response"`
	JSON        string   `json:"json,omitempty" xml:"json,omitempty" yaml:"json,omitempty"`
	XML         string   `json:"xml,omitempty" xml:"xml,omitempty" yaml:"xml,omitempty"`
	YAML        string   `json:"yaml,omitempty" xml:"yaml,omitempty" yaml:"yaml,omitempty"`
}

// XMLString method
func (er ExampleResponse) XMLString() (result string) {
	if b, err := xml.MarshalIndent(er.Object, "", "    "); err == nil {
		result = xml.Header + string(b)
	}
	return
}

// JSONString method
func (er ExampleResponse) JSONString() (result string) {
	if b, err := json.MarshalIndent(er.Object, "", "    "); err == nil {
		result = string(b)
	}
	return
}

// YAMLString method
func (er ExampleResponse) YAMLString() (result string) {
	if b, err := yaml.Marshal(er.Object); err == nil {
		result = string(b)
	}
	return
}

// ExampleResponses is array of Command
type ExampleResponses []ExampleResponse

// MarshalXML for marshal into XML
func (arr ExampleResponses) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	a := []interface{}{}
	for _, item := range arr {
		a = append(a, item)
	}
	return MarshallerXML(a)(e, start)
}
