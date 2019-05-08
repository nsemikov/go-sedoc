package sedoc

import (
	"encoding/xml"
	"fmt"
)

// NewError is Error constructor
func (api API) NewError(code int, details ...interface{}) *Error {
	return api.NewErrorInternal(code, nil, details...)
}

// NewErrorInternal is Error constructor
func (api API) NewErrorInternal(code int, internal error, details ...interface{}) *Error {
	err := api.Errors.Get(code)
	err.Internal = internal
	if len(details) > 0 {
		err.Description = err.Description + ":"
		for _, detail := range details {
			err.Description = err.Description + " " + fmt.Sprint(detail)
		}
	}
	return &err
}

// Error is standard quickq error type
type Error struct {
	XMLName     xml.Name `json:"-" xml:"error" yaml:"-"`
	Code        int      `json:"code" xml:"code,attr" yaml:"code"`
	Description string   `json:"desc" xml:"desc,attr" yaml:"desc"`
	Internal    error    `json:"-" xml:"-" yaml:"-"`
}

// Error convert to string
func (e Error) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("[%d] %s |internal| %v", e.Code, e.Description, e.Internal)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Description)
}

const (
	// ErrUnknown means an unknown error occurred
	ErrUnknown = iota + 1
	// ErrInvalidRequest means catched request is invalid
	ErrInvalidRequest
	// ErrUnknownCommand means catched unknown command
	ErrUnknownCommand
	// ErrInvalidArgumentRegExp means invalid command argument parameter regexp
	ErrInvalidArgumentRegExp
	// ErrArgumentRegExpMatchFails means match command argument parameter regexp fails
	ErrArgumentRegExpMatchFails
	// ErrRequiredArgumentMissing means require command argument parameter missing
	ErrRequiredArgumentMissing
	// ErrUnknownArgument means unknown command argument parameter in request
	ErrUnknownArgument
	// ErrInvalidArgumentValue means invalid command argument parameter value
	ErrInvalidArgumentValue
	// LastUsedErrorCode is last error code used in sedoc
	LastUsedErrorCode = 100
)

// DefaultErrors is map of Mrror
var DefaultErrors = Errors{
	{Code: ErrUnknown, Description: "unknown error occurred"},
	{Code: ErrInvalidRequest, Description: "can't parse request"},
	{Code: ErrUnknownCommand, Description: "unknown command"},
	{Code: ErrInvalidArgumentRegExp, Description: "invalid command argument parameter regexp"},
	{Code: ErrArgumentRegExpMatchFails, Description: "match command argument parameter regexp fails"},
	{Code: ErrRequiredArgumentMissing, Description: "require command argument parameter missing"},
	{Code: ErrUnknownArgument, Description: "unknown command argument parameter in request"},
	{Code: ErrInvalidArgumentValue, Description: "invalid command argument parameter value"},
}

// Errors is array of Error
type Errors []Error

// Contains Error
func (arr *Errors) Contains(code int) bool {
	for _, err := range *arr {
		if err.Code == code {
			return true
		}
	}
	return false
}

// Add is add error to array
func (arr *Errors) Add(err Error) {
	if arr.Contains(err.Code) {
		panic("api: duplicate error")
	}
	*arr = append(*arr, err)
}

// Get is get error from array
func (arr *Errors) Get(code int) Error {
	for _, err := range *arr {
		if err.Code == code {
			return err
		}
	}
	return Error{
		Code: code,
	}
}

// Remove is remove error from array
func (arr *Errors) Remove(code int) {
	for idx, err := range *arr {
		if err.Code == code {
			*arr = append(append(Errors{}, (*arr)[:idx]...), (*arr)[idx+1:]...)
			return
		}
	}
	panic("api: unknown error")
}

// MarshalXML for marshal into XML
func (arr Errors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
