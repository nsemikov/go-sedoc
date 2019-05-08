package sedoc_test

import (
	"errors"

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

// Use Set for arguments that must be setted to new or exist items
func Example_setArguments() {
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
func Example_commandExamples() {
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

func Example_middleware() {
	// ...
	api.Use(func(next sedoc.HandlerFunc) sedoc.HandlerFunc {
		return func(c sedoc.Context) error {
			// your code here
			return next(c)
		}
	})
	// ...
}

func Example_customContext() {
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

func Example_customError() {
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
