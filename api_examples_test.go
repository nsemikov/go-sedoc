package sedoc_test

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"time"

	sedoc "github.com/stdatiks/go-sedoc"
	"github.com/stdatiks/go-sedoc/argument"
)

// Here is basic example of sedoc api usage
func Example_basic() {
	// ...
	a := sedoc.New()
	a.Description = "My Service API"
	a.AddCommand(sedoc.Command{
		Name:        "info",
		Description: "Get information about service.",
		Handler: func(c sedoc.Context) error {
			c.Response().Result = &struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Version     string `json:"version"`
			}{
				Name:        "mysrv",
				Description: "My Service",
				Version:     "1.0.0",
			}
			return nil
		},
	})
	// ...
}

var api = sedoc.New()

func Example_advanced() {
	// create api
	api := sedoc.New()
	api.Description = "My self documented API"
	// create commands
	infoCmd := sedoc.Command{
		Name: "info",
		Handler: func(c sedoc.Context) error {
			c.Response().Result = fmt.Sprintf("%s v%s", "mysrv", "1.0.0")
			return nil
		},
	}
	// add examples to command (optionaly)
	infoCmd.Examples = sedoc.Examples{
		sedoc.Example{
			Name: "simple",
			Request: sedoc.ExampleRequest{
				Object: sedoc.Request{
					Datetime: func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-10-16T09:58:03.487508407Z"); return t }(),
					Command:  "info",
				},
			},
			Responses: sedoc.ExampleResponses{
				sedoc.ExampleResponse{
					Name: "simple",
					Object: sedoc.Response{
						Datetime: func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-10-16T09:58:03.487508407Z"); return t }(),
						Command:  "info",
						Result:   "MyService v0.1.0",
					},
				},
			},
		},
	}
	// add command to api
	api.AddCommand(infoCmd)
	// create request
	request := sedoc.NewRequest()
	request.Command = "help"
	// and execute it
	response := api.Execute(request)
	// at last marshal response to XML (currently supported are XML, JSON and YAML)
	if b, err := xml.MarshalIndent(response, "", "  "); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}
}

// Use Set for arguments that must be setted to new or exist items
func Example_set() {
	// ...
	api.AddCommand(sedoc.Command{
		Name:        "user.add",
		Description: "Create new user (by another one).",
		Handler: func(c sedoc.Context) error {
			// ...
			return nil
		},
		Set: sedoc.Arguments{
			sedoc.Argument{
				Name:        "login",
				Type:        sedoc.ArgumentTypeString,
				Description: "Login string.",
				Required:    true,
			},
			sedoc.Argument{
				Name:        "password",
				Type:        sedoc.ArgumentTypeString,
				Description: "Password string.",
				Required:    true,
			},
			argument.Email, // you can use some args from "github.com/stdatiks/go-sedoc/argument"
			argument.Name,  // you can use some args from "github.com/stdatiks/go-sedoc/argument"
		},
	})
	// ...
}

// Use Where for arguments which must be used to find exist items
func Example_where() {
	// ...
	api.AddCommand(sedoc.Command{
		Name:        "user.get",
		Description: "Get existed single user or user list.",
		Handler: func(c sedoc.Context) error {
			// ...
			return nil
		},
		Where: sedoc.Arguments{
			argument.Required(argument.Count),
		},
	})
	// ...
}

// Use Arguments for other request parameters
func Example_arguments() {
	// ...
	api.AddCommand(sedoc.Command{
		Name:        "user.get",
		Description: "Get existed single user or user list.",
		Handler: func(c sedoc.Context) error {
			// ...
			return nil
		},
		Arguments: sedoc.Arguments{
			argument.Count,
			argument.Offset,
		},
	})
	// ...
}

// Use Examples to show how to use your API
func ExampleExample_usage() {
	// ...
	api.AddCommand(sedoc.Command{
		Name:        "signin",
		Description: "Sign in (login). Create new session.",
		Where: sedoc.Arguments{
			argument.Required(argument.Login),
			argument.Required(argument.Password),
		},
		Handler: func(c sedoc.Context) error {
			// ...
			return nil
		},
		Examples: []sedoc.Example{
			sedoc.Example{
				Name:        "valid",
				Description: "valid signin request",
				Request: sedoc.ExampleRequest{
					Object: *sedoc.NewRequest(), // Your request example here
				},
				Responses: []sedoc.ExampleResponse{
					sedoc.ExampleResponse{
						Object: *sedoc.NewResponse(), // Your response example here
					},
					sedoc.ExampleResponse{
						Description: "user not found or password incorrect",
						Object:      *sedoc.NewResponse(), // Your response example here
					},
				},
			},
			sedoc.Example{
				Name:        "invalid",
				Description: "invalid signin request (password is not present)",
				Request: sedoc.ExampleRequest{
					Object: *sedoc.NewRequest(), // Your request example here
				},
				Responses: []sedoc.ExampleResponse{
					sedoc.ExampleResponse{
						Object: *sedoc.NewResponse(), // Your response example here
					},
				},
			},
		},
	})
	// ...
}

func ExampleMiddlewareFunc_usage() {
	// ...
	api.Use(func(next sedoc.HandlerFunc) sedoc.HandlerFunc {
		return func(c sedoc.Context) error {
			// your code here
			return next(c)
		}
	})
	// ...
}

func ExampleContext_custom() {
	// ...
	type mycontext struct {
		sedoc.Context
	}
	// ...
	api.Use(func(next sedoc.HandlerFunc) sedoc.HandlerFunc {
		return func(c sedoc.Context) (err error) {
			return next(&mycontext{c})
		}
	})
	// ...
	api.AddCommand(sedoc.Command{
		Name: "signin",
		Handler: func(c sedoc.Context) error {
			cc, ok := c.(*mycontext)
			if !ok || cc == nil {
				return errors.New("invalid context")
			}
			// ...
			return nil
		},
		Set: sedoc.Arguments{
			argument.Required(argument.Login),
			argument.Required(argument.Login),
		},
	})
	// ...
}

func ExampleErrors_custom() {
	// ...
	const (
		ErrCommandNotImplemented = iota + sedoc.LastUsedErrorCode
		ErrAuthRequired
		ErrAuth
	)
	var ErrorMap = sedoc.Errors{
		{Code: ErrCommandNotImplemented, Description: "Command not implemented yet"},
		{Code: ErrAuthRequired, Description: "No credentials set"},
		{Code: ErrAuth, Description: "Auth error"},
	}
	// ...
	api.Errors = append(sedoc.DefaultErrors, ErrorMap...)
	// ...
	api.AddCommand(sedoc.Command{
		Name: "exec",
		Handler: func(c sedoc.Context) error {
			// ...
			return c.Error(ErrCommandNotImplemented, "exec")
		},
	})
	// ...
}
