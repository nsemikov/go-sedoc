package sedoc

import (
	"time"
)

// Context interface
type Context interface {
	// Request return instance of Request
	Request() *Request
	// Response return instance of Response
	Response() *Response
	// Command return copy of Command
	Command() *Command
	// Error method
	Error(code int, details ...interface{}) *Error
	// ErrorInternal method
	ErrorInternal(code int, internal error, details ...interface{}) *Error
}

type context struct {
	api  *API
	req  *Request
	resp *Response
	cmd  *Command
}

// Request return instance of Request
func (c *context) Request() *Request {
	if c.req == nil {
		c.req = NewRequest()
	}
	return c.req
}

// Response return instance of Response
func (c *context) Response() *Response {
	if c.resp == nil {
		c.resp = NewResponse()
	}
	return c.resp
}

// Command return copy of Command
func (c *context) Command() *Command {
	if c.cmd == nil {
		if c.api == nil {
			panic("sedoc: Context can`t contain nil API")
		}
		c.cmd = &Command{}
		*c.cmd = c.api.GetCommand(c.Request().Command)
	}
	return c.cmd
}

// Error method
func (c context) Error(code int, details ...interface{}) *Error {
	return c.api.NewError(code, details...)
}

// ErrorInternal method
func (c context) ErrorInternal(code int, internal error, details ...interface{}) *Error {
	return c.api.NewErrorInternal(code, internal, details...)
}

func (c *context) checkRequestArguments() (err error) {
	return c.Command().checkArguments(c.Request())
}

func fillResponseMissingDataFromRequest(request *Request, response *Response) {
	if request == nil || response == nil {
		return
	}
	if len(request.ID) > 0 {
		response.ID = request.ID
	}
	if !request.Datetime.IsZero() && response.Datetime.IsZero() {
		response.Datetime = time.Now().UTC()
	}
	if len(response.Command) <= 0 {
		response.Command = request.Command
	}
	if len(response.Session) <= 0 {
		response.Session = request.Session
	}
	if response.Error != nil && response.Error.Code == 0 {
		response.Error = nil
	}
}
