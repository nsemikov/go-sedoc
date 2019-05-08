package sedoc

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

// Request contains request fields
type Request struct {
	XMLName   xml.Name       `json:"-" xml:"request" yaml:"-"`
	ID        string         `json:"id,omitempty" xml:"id,attr,omitempty" yaml:"id,omitempty"`
	Datetime  time.Time      `json:"datetime,omitempty" xml:"datetime,attr,omitempty" yaml:"datetime,omitempty"`
	Session   string         `json:"session,omitempty" xml:"session,attr,omitempty" yaml:"session,omitempty"`
	Command   string         `json:"command,omitempty" xml:"command,attr" yaml:"command,omitempty"`
	Arguments InterfaceMap   `json:"args,omitempty" xml:"args,omitempty" yaml:"args,omitempty"`
	Where     []InterfaceMap `json:"where,omitempty" xml:"where,omitempty" yaml:"where,omitempty"`
	Set       InterfaceMap   `json:"set,omitempty" xml:"set,omitempty" yaml:"set,omitempty"`
}

// String format Request as string
func (r *Request) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// NewRequest func
func NewRequest() *Request {
	return &Request{
		Arguments: make(InterfaceMap),
		Where:     make([]InterfaceMap, 0),
		Set:       make(InterfaceMap),
	}
}
