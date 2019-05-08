package sedoc

import (
	"fmt"
	"time"
)

// API struct
type API struct {
	Description     string                 `json:"description,omitempty" xml:"description,attr,omitempty" yaml:"description,omitempty"`
	RequestFormat   Arguments              `json:"request_format,omitempty" xml:"request_format,omitempty"  yaml:"request_format,omitempty"`
	ResponseFormat  Arguments              `json:"response_format,omitempty" xml:"response_format,omitempty"  yaml:"response_format,omitempty"`
	Commands        Commands               `json:"commands,omitempty" xml:"commands,omitempty" yaml:"commands,omitempty"`
	Errors          Errors                 `json:"errors,omitempty" xml:"errors,omitempty" yaml:"errors,omitempty"`
	PrefixArguments string                 `json:"-" xml:"-" yaml:"-"`
	PrefixSet       string                 `json:"-" xml:"-" yaml:"-"`
	PrefixWhere     string                 `json:"-" xml:"-" yaml:"-"`
	ErrorHandler    CustomErrorHandlerFunc `json:"-" xml:"-" yaml:"-"`
	middleware      []MiddlewareFunc
}

// HandlerFunc func
type HandlerFunc func(Context) error

// MiddlewareFunc func
type MiddlewareFunc func(HandlerFunc) HandlerFunc

// CustomErrorHandlerFunc func
type CustomErrorHandlerFunc func(error, Context)

// New api constructor
func New() (api *API) {
	id := Argument{
		Name:        "id",
		Type:        ArgumentTypeString,
		Description: "Request identifier (for debugging)",
	}
	datetime := Argument{
		Name:        "datetime",
		Type:        ArgumentTypeString,
		Description: "Datetime string (ISO 8601)",
	}
	session := Argument{
		Name:        "session",
		Type:        ArgumentTypeUUID,
		Description: "Session token uuid, formatted like \"01234567-89ab-cdef-0123-456789abcdef\"",
	}
	command := Argument{
		Name:        "command",
		Type:        ArgumentTypeString,
		Required:    true,
		Description: "Command name string",
	}
	args := Argument{
		Name:        "args",
		Type:        argumentTypeObject,
		Description: "Extra request parameters, one-level object",
	}
	where := Argument{
		Name:        "where",
		Type:        argumentTypeArray,
		Description: "Search item(s) parameters, simple Array of one-level objects",
	}
	set := Argument{
		Name:        "set",
		Type:        argumentTypeObject,
		Description: "Item(s) data to set, one-level object",
	}
	result := Argument{
		Name:        "result",
		Type:        argumentTypeObject,
		Description: "Result object. For XML maybe used another name",
	}
	errorArg := Argument{
		Name:        "error",
		Type:        argumentTypeObject,
		Description: "Error object. Contains `code` and `desc` fields",
	}
	api = &API{
		Errors:   DefaultErrors,
		Commands: Commands{},
		RequestFormat: Arguments{
			id,
			datetime,
			session,
			command,
			args,
			where,
			set,
		},
		ResponseFormat: Arguments{
			id,
			datetime,
			session,
			command,
			args,
			result,
			errorArg,
		},
		PrefixArguments: "",
		PrefixSet:       "2-",
		PrefixWhere:     "4-",
	}
	api.ErrorHandler = func(err error, c Context) {
		if err != nil {
			var ok bool
			c.Response().Error, ok = err.(*Error)
			if !ok {
				serr := DefaultErrors.Get(ErrUnknown)
				c.Response().Error = &serr
				c.Response().Error.Internal = err
			}
		}
	}

	api.AddCommand(Command{
		Name:        "help",
		Description: "Get list of commands",
		Handler: func(c Context) error {
			for cidx := range api.Commands {
				command := &api.Commands[cidx]
				for eidx := range command.Examples {
					example := &command.Examples[eidx]
					example.Request.JSON = example.Request.JSONString()
					example.Request.XML = example.Request.XMLString()
					example.Request.YAML = example.Request.YAMLString()
					for ridx := range example.Responses {
						response := &example.Responses[ridx]
						response.JSON = response.JSONString()
						response.XML = response.XMLString()
						response.YAML = response.YAMLString()
					}
				}
			}
			c.Response().Result = api
			return nil
		},
		Examples: []Example{
			{
				Name:        "simple help",
				Description: "simple help command usage example",
				Request: ExampleRequest{
					Object: Request{
						Datetime: func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-10-16T09:58:03.487508407Z"); return t }(),
						Command:  "help",
					},
				},
			},
		},
	})
	return
}

// Use add middleware
func (api *API) Use(middleware ...MiddlewareFunc) {
	if api.middleware == nil {
		api.middleware = []MiddlewareFunc{}
	}
	api.middleware = append(api.middleware, middleware...)
}

// GetCommand return Command
func (api *API) GetCommand(name string) Command {
	if cmd, err := api.Commands.Get(name); err == nil {
		return cmd
	}
	return Command{}
}

// AddCommand append handler with help to API.
// AddCommand will rewrite command with the same name
func (api *API) AddCommand(cmd Command) {
	api.RemoveCommand(cmd.Name)
	_ = api.Commands.Add(cmd)
}

// RemoveCommand remove command from API
func (api *API) RemoveCommand(name string) {
	if _, err := api.Commands.Get(name); err != nil {
		_ = api.Commands.Remove(name)
	}
}

// Execute command from API
func (api *API) Execute(request *Request) *Response {
	c := &context{api: api, req: request}
	var err error
	if c.req == nil {
		err = api.NewError(ErrInvalidRequest)
	} else if !c.Command().Valid() {
		err = api.NewError(ErrUnknownCommand)
	} else {
		err = c.checkRequestArguments()
	}
	if err == nil {
		handler := c.Command().Handler
		for idx := len(api.middleware) - 1; idx >= 0; idx-- {
			handler = api.middleware[idx](handler)
		}
		err = handler(c)
	}
	if err != nil {
		serr, ok := err.(*Error)
		if !ok {
			serr = api.NewError(ErrUnknown)
			serr.Description = fmt.Sprintf("%s: %s", serr.Description, err.Error())
		}
		c.api.ErrorHandler(serr, c)
	}
	fillResponseMissingDataFromRequest(c.Request(), c.Response())
	return c.Response()
}
