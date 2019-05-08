package sedoc

import "testing"

func TestCommand_checkArguments(t *testing.T) {
	type fields struct {
		Name            string
		Description     string
		AuthRequired    bool
		SessionRequired bool
		ResponseList    bool
		Arguments       Arguments
		Where           Arguments
		Set             Arguments
		Handler         HandlerFunc
		Examples        []Example
	}
	f := fields{
		Arguments: Arguments{
			Argument{Name: "req", Type: ArgumentTypeInteger, RegExp: "^\\d+$", Required: true},
		},
		Set: Arguments{
			Argument{Name: "name", Type: ArgumentTypeString, RegExp: ""},
			Argument{Name: "alias", Type: ArgumentTypeString, RegExp: "^.{0,8}$"},
		},
		Where: Arguments{
			Argument{Name: "id", Type: ArgumentTypeInteger, RegExp: "^\\d$"},
			Argument{Name: "name", Type: ArgumentTypeString, RegExp: ""},
			Argument{Name: "alias", Type: ArgumentTypeString, RegExp: "^.({8}$"},
		},
	}
	type args struct {
		request *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"params.req:nil", f, args{request: NewRequest()}, true},
		{"params.req:nil", f, args{request: &Request{
			Arguments: InterfaceMap{},
			Set:       InterfaceMap{},
			Where:     []InterfaceMap{},
		}}, true},
		{"params.req:invalid", f, args{request: &Request{
			Arguments: InterfaceMap{"req": "10a"},
		}}, true},
		{"params.unknown", f, args{request: &Request{
			Arguments: InterfaceMap{"req": "10", "unknown": false},
		}}, true},
		{"set.alias:unmatch", f, args{request: &Request{
			Arguments: InterfaceMap{"req": "10"},
			Set:       InterfaceMap{"alias": "0123456789"},
		}}, true},
		{"set.alias:good", f, args{request: &Request{
			Arguments: InterfaceMap{"req": "10"},
			Set:       InterfaceMap{"alias": "xmpl"},
		}}, false},
		{"where.alias:unmatch", f, args{request: &Request{
			Arguments: InterfaceMap{"req": "10"},
			Set:       InterfaceMap{"alias": "xmpl"},
			Where: []InterfaceMap{
				{"alias": "xmpl"},
			},
		}}, true},
		{"all_good", f, args{request: &Request{
			Arguments: InterfaceMap{"req": "10"},
			Set:       InterfaceMap{"alias": "xmpl"},
			Where: []InterfaceMap{
				{"name": "some name =)"},
			},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &Command{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Arguments:   tt.fields.Arguments,
				Where:       tt.fields.Where,
				Set:         tt.fields.Set,
				Handler:     tt.fields.Handler,
				Examples:    tt.fields.Examples,
			}
			if err := cmd.checkArguments(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Command.checkArguments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
