package sedoc

import (
	"encoding/xml"
	"fmt"
	"regexp"
)

// Argument is api command argument
type Argument struct {
	XMLName     xml.Name     `json:"-" xml:"arg" yaml:"-"`
	Name        string       `json:"name" xml:"name,attr" yaml:"name"`
	Type        ArgumentType `json:"type" xml:"type,attr" yaml:"type"`
	Description string       `json:"description" xml:"description,attr" yaml:"description"`
	Nullable    bool         `json:"nullable,omitempty" xml:"nullable,attr,omitempty" yaml:"nullable,omitempty"`
	Multiple    bool         `json:"multiple,omitempty" xml:"multiple,attr,omitempty" yaml:"multiple,omitempty"`
	Required    bool         `json:"required,omitempty" xml:"required,attr,omitempty" yaml:"required,omitempty"`
	Disabled    bool         `json:"-" xml:"-" yaml:"-"`
	RegExp      string       `json:"regexp,omitempty" xml:"regexp,attr,omitempty" yaml:"regexp,omitempty"`
}

// Arguments is array of Argument
type Arguments []Argument

// Contains Argument
func (arr *Arguments) Contains(name string) bool {
	_, err := arr.Get(name)
	return err == nil
}

// Add is add argument to array
func (arr *Arguments) Add(arg Argument) error {
	if arr.Contains(arg.Name) {
		panic("argument already exist: " + arg.Name)
	}
	*arr = append(*arr, arg)
	return nil
}

// Get is get argument from array
func (arr *Arguments) Get(name string) (Argument, error) {
	for _, arg := range *arr {
		if arg.Name == name && !arg.Disabled {
			return arg, nil
		}
	}
	return Argument{}, DefaultErrors.Get(ErrUnknownArgument)
}

// Remove is remove argument from array
func (arr *Arguments) Remove(name string) error {
	for idx, arg := range *arr {
		if arg.Name == name {
			*arr = append(append(Arguments{}, (*arr)[:idx]...), (*arr)[idx+1:]...)
			return nil
		}
	}
	return DefaultErrors.Get(ErrUnknownArgument)
}

// MarshalXML for marshal into XML
func (arr Arguments) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	required := 0
	a := []interface{}{}
	for _, item := range arr {
		a = append(a, item)
		if item.Required {
			required++
		}
	}
	if required > 0 {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Space: "", Local: "required"},
			Value: fmt.Sprint(required),
		})
	}
	return MarshallerXML(a)(e, start)
}

// Match func is matching stringify representation of v for Argument RegExp
func (arg Argument) Match(v interface{}) (bool, error) {
	res, err := regexp.MatchString(arg.RegExp, fmt.Sprintf("%v", v))
	return res, err
}

// ArgumentType is enum of supported types of Argument
type ArgumentType string

const (
	// ArgumentTypeBoolean is boolean Argument type
	ArgumentTypeBoolean ArgumentType = "boolean"
	// ArgumentTypeInteger is integer Argument type
	ArgumentTypeInteger ArgumentType = "integer"
	// ArgumentTypeFloat is integer Argument type
	ArgumentTypeFloat ArgumentType = "float"
	// ArgumentTypeString is string Argument type
	ArgumentTypeString ArgumentType = "string"
	// ArgumentTypeDuration is duration Argument type
	ArgumentTypeDuration ArgumentType = "duration"
	// ArgumentTypeUUID is duration Argument type
	ArgumentTypeUUID ArgumentType = "uuid"
	// ArgumentTypeTime is datetime Argument type
	ArgumentTypeTime ArgumentType = "datetime"

	argumentTypeObject ArgumentType = "object"
	argumentTypeArray  ArgumentType = "array"

	// ArgumentTypeList is duration Argument type
	ArgumentTypeList ArgumentType = "list"
)

// String is string convertor for ArgumentType
func (t ArgumentType) String() string {
	return string(t)
}
