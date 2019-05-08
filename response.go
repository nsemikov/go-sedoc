package sedoc

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

// Response contains response fields
type Response struct {
	XMLName   xml.Name     `json:"-" xml:"response" yaml:"-"`
	ID        string       `json:"id,omitempty" xml:"id,attr,omitempty" yaml:"id,omitempty"`
	Datetime  time.Time    `json:"datetime,omitempty" xml:"datetime,attr,omitempty" yaml:"datetime,omitempty"`
	Session   string       `json:"session,omitempty" xml:"session,attr,omitempty" yaml:"session,omitempty"`
	Command   string       `json:"command,omitempty" xml:"command,attr" yaml:"command,omitempty"`
	Arguments InterfaceMap `json:"args,omitempty" xml:"args,omitempty" yaml:"args,omitempty"`
	Result    interface{}  `json:"result,omitempty" xml:"result,omitempty" yaml:"result,omitempty"`
	Error     *Error       `json:"error,omitempty" xml:"error,omitempty" yaml:"error,omitempty"`
}

// String format Response as string
func (r *Response) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

// NewResponse func
func NewResponse() *Response {
	return &Response{
		Error: &Error{},
	}
}
