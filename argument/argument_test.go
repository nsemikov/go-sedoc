package argument

import (
	"reflect"
	"testing"

	sedoc "github.com/stdatiks/go-sedoc"
)

func TestRequired(t *testing.T) {
	type args struct {
		arg sedoc.Argument
	}
	tests := []struct {
		name string
		args args
		want sedoc.Argument
	}{
		// TODO: Add test cases
		{"", args{sedoc.Argument{}}, sedoc.Argument{Required: true}},
		{"", args{sedoc.Argument{Name: "test"}}, sedoc.Argument{Name: "test", Required: true}},
		{"", args{sedoc.Argument{Name: "test", Required: false}}, sedoc.Argument{Name: "test", Required: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Required(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Required() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullable(t *testing.T) {
	type args struct {
		arg sedoc.Argument
	}
	tests := []struct {
		name string
		args args
		want sedoc.Argument
	}{
		// TODO: Add test cases
		{"", args{sedoc.Argument{}}, sedoc.Argument{Nullable: true}},
		{"", args{sedoc.Argument{Name: "test"}}, sedoc.Argument{Name: "test", Nullable: true}},
		{"", args{sedoc.Argument{Name: "test", Nullable: false}}, sedoc.Argument{Name: "test", Nullable: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Nullable(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Nullable() = %v, want %v", got, tt.want)
			}
		})
	}
}
