package sedoc

import (
	"reflect"
	"testing"
)

func Test_context_Request(t *testing.T) {
	type fields struct {
		api  *API
		req  *Request
		resp *Response
		cmd  *Command
	}
	tests := []struct {
		name   string
		fields fields
		want   *Request
	}{
		// TODO: Add test cases.
		{"", fields{nil, nil, nil, nil}, &Request{
			Arguments: make(InterfaceMap),
			Where:     make([]InterfaceMap, 0),
			Set:       make(InterfaceMap),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &context{
				api:  tt.fields.api,
				req:  tt.fields.req,
				resp: tt.fields.resp,
				cmd:  tt.fields.cmd,
			}
			if got := c.Request(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("context.Request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_context_Response(t *testing.T) {
	type fields struct {
		api  *API
		req  *Request
		resp *Response
		cmd  *Command
	}
	tests := []struct {
		name   string
		fields fields
		want   *Response
	}{
		// TODO: Add test cases.
		{"", fields{nil, nil, nil, nil}, &Response{Error: &Error{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &context{
				api:  tt.fields.api,
				req:  tt.fields.req,
				resp: tt.fields.resp,
				cmd:  tt.fields.cmd,
			}
			if got := c.Response(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("context.Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_context_Command(t *testing.T) {
	a := New()
	a.AddCommand(Command{
		Name: "test1",
	})
	a.AddCommand(Command{
		Name:        "test2",
		Description: "with desc",
	})
	type fields struct {
		api  *API
		req  *Request
		resp *Response
		cmd  *Command
	}
	tests := []struct {
		name      string
		fields    fields
		want      *Command
		wantPanic bool
	}{
		// TODO: Add test cases.
		{"", fields{nil, nil, nil, nil}, &Command{}, true},
		{"", fields{a, &Request{Command: "test1"}, nil, nil}, &Command{Name: "test1"}, false},
		{"", fields{a, &Request{Command: "test2"}, nil, nil}, &Command{Name: "test2", Description: "with desc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				p := recover()
				if tt.wantPanic && p == nil {
					t.Errorf("context.Command() want panic, but not")
				} else if !tt.wantPanic && p != nil {
					t.Errorf("context.Command() panic: %v", p)
				}
			}()
			c := &context{
				api:  tt.fields.api,
				req:  tt.fields.req,
				resp: tt.fields.resp,
				cmd:  tt.fields.cmd,
			}
			if got := c.Command(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("context.Command() = %v, want %v", got, tt.want)
			}
		})
	}
}
