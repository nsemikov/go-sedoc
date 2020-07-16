package sedoc

import (
	"reflect"
	"testing"
	"time"

	"github.com/nsemikov/go-sedoc/types"
)

func TestArgumentType_Parse(t *testing.T) {
	var incompatibleType ArgumentType = "incompatible_type"
	var nilParser ArgumentType = "nil_parser"
	SetArgumentParser(nilParser, nil)
	type args struct {
		v    interface{}
		list bool
	}
	tests := []struct {
		name    string
		t       ArgumentType
		args    args
		wantR   interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{`incompatible_type`, incompatibleType, args{nil, false}, nil, true},
		{`nil_parser`, nilParser, args{nil, false}, nil, true},

		{`incompatible_boolean_type`, ArgumentTypeBoolean, args{nil, false}, false, true},
		{`boolean(false)`, ArgumentTypeBoolean, args{false, false}, false, false},
		{`boolean(true)`, ArgumentTypeBoolean, args{true, false}, true, false},
		{`boolean(0)`, ArgumentTypeBoolean, args{0, false}, false, false},
		{`boolean(1)`, ArgumentTypeBoolean, args{1, false}, true, false},
		{`boolean(-1)`, ArgumentTypeBoolean, args{-1, false}, true, false},
		{`boolean(15.2)`, ArgumentTypeBoolean, args{15.2, false}, true, false},
		{`boolean(-8)`, ArgumentTypeBoolean, args{-8, false}, true, false},
		{`boolean("-5.8")`, ArgumentTypeBoolean, args{"-5.8", false}, false, false},
		{`boolean("false")`, ArgumentTypeBoolean, args{"false", false}, false, false},
		{`boolean("true")`, ArgumentTypeBoolean, args{"true", false}, true, false},
		{`boolean("foo")`, ArgumentTypeBoolean, args{"foo", false}, false, false},
		{`boolean("2s")`, ArgumentTypeBoolean, args{"2s", false}, false, false},
		{`boolean("1.5h25s")`, ArgumentTypeBoolean, args{"1.5h25s", false}, false, false},
		{`boolean("-15m8s")`, ArgumentTypeBoolean, args{"-15m8s", false}, false, false},

		{`incompatible_integer_type`, ArgumentTypeInteger, args{nil, false}, 0, true},
		{`integer(false)`, ArgumentTypeInteger, args{false, false}, 0, false},
		{`integer(true)`, ArgumentTypeInteger, args{true, false}, 1, false},
		{`integer(0)`, ArgumentTypeInteger, args{0, false}, 0, false},
		{`integer(1)`, ArgumentTypeInteger, args{1, false}, 1, false},
		{`integer(-1)`, ArgumentTypeInteger, args{-1, false}, -1, false},
		{`integer(15.2)`, ArgumentTypeInteger, args{15.2, false}, 15, false},
		{`integer(-8)`, ArgumentTypeInteger, args{-8, false}, -8, false},
		{`integer("-5.8")`, ArgumentTypeInteger, args{"-5.8", false}, 0, true},
		{`integer("false")`, ArgumentTypeInteger, args{"false", false}, 0, true},
		{`integer("true")`, ArgumentTypeInteger, args{"true", false}, 0, true},
		{`integer("foo")`, ArgumentTypeInteger, args{"foo", false}, 0, true},
		{`integer("2s")`, ArgumentTypeInteger, args{"2s", false}, 0, true},
		{`integer("1.5h25s")`, ArgumentTypeInteger, args{"1.5h25s", false}, 0, true},
		{`integer("-15m8s")`, ArgumentTypeInteger, args{"-15m8s", false}, 0, true},

		{`incompatible_float_type`, ArgumentTypeFloat, args{nil, false}, 0.0, true},
		{`float(false)`, ArgumentTypeFloat, args{false, false}, 0.0, false},
		{`float(true)`, ArgumentTypeFloat, args{true, false}, 1.0, false},
		{`float(0)`, ArgumentTypeFloat, args{0, false}, 0.0, false},
		{`float(1)`, ArgumentTypeFloat, args{1, false}, 1.0, false},
		{`float(-1)`, ArgumentTypeFloat, args{-1, false}, -1.0, false},
		{`float(15.2)`, ArgumentTypeFloat, args{15.2, false}, 15.2, false},
		{`float(-8)`, ArgumentTypeFloat, args{-8, false}, -8.0, false},
		{`float("-5.8")`, ArgumentTypeFloat, args{"-5.8", false}, -5.8, false},
		{`float("false")`, ArgumentTypeFloat, args{"false", false}, 0.0, true},
		{`float("true")`, ArgumentTypeFloat, args{"true", false}, 0.0, true},
		{`float("foo")`, ArgumentTypeFloat, args{"foo", false}, 0.0, true},
		{`float("2s")`, ArgumentTypeFloat, args{"2s", false}, 0.0, true},
		{`float("1.5h25s")`, ArgumentTypeFloat, args{"1.5h25s", false}, 0.0, true},
		{`float("-15m8s")`, ArgumentTypeFloat, args{"-15m8s", false}, 0.0, true},

		{`duration(false)`, ArgumentTypeDuration, args{false, false}, types.Duration(0), true},
		{`duration(true)`, ArgumentTypeDuration, args{true, false}, types.Duration(0), true},
		{`duration(0)`, ArgumentTypeDuration, args{0, false}, types.Duration(0), false},
		{`duration(1)`, ArgumentTypeDuration, args{1, false}, types.Duration(1 * time.Nanosecond), false},
		{`duration(-1)`, ArgumentTypeDuration, args{-1, false}, types.Duration(-1 * time.Nanosecond), false},
		{`duration(15.2)`, ArgumentTypeDuration, args{15.2, false}, types.Duration(0), true},
		{`duration(-8)`, ArgumentTypeDuration, args{-8, false}, types.Duration(-8 * time.Nanosecond), false},
		{`duration("-5.8")`, ArgumentTypeDuration, args{"-5.8", false}, types.Duration(0), true},
		{`duration("false")`, ArgumentTypeDuration, args{"false", false}, types.Duration(0), true},
		{`duration("true")`, ArgumentTypeDuration, args{"true", false}, types.Duration(0), true},
		{`duration("foo")`, ArgumentTypeDuration, args{"foo", false}, types.Duration(0), true},
		{`duration("2s")`, ArgumentTypeDuration, args{"2s", false}, types.Duration(2 * time.Second), false},
		{`duration("1.5h25s")`, ArgumentTypeDuration, args{"1.5h25s", false}, types.Duration(1*time.Hour + 30*time.Minute + 25*time.Second), false},
		{`duration("-15m8s")`, ArgumentTypeDuration, args{"-15m8s", false}, types.Duration(-15*time.Minute - 8*time.Second), false},
		{`duration(time.Duration:1m")`, ArgumentTypeDuration, args{time.Duration(1 * time.Minute), false}, types.Duration(1 * time.Minute), false},
		{`duration(types.Duration:15m")`, ArgumentTypeDuration, args{types.Duration(15 * time.Minute), false}, types.Duration(15 * time.Minute), false},

		{`incompatible_string_type`, ArgumentTypeString, args{nil, false}, "", true},
		{`string(false)`, ArgumentTypeString, args{false, false}, "false", false},
		{`string(true)`, ArgumentTypeString, args{true, false}, "true", false},
		{`string(0)`, ArgumentTypeString, args{0, false}, "0", false},
		{`string(0.0)`, ArgumentTypeString, args{float64(0.0), false}, "0.0", false},
		{`string(1)`, ArgumentTypeString, args{1, false}, "1", false},
		{`string(-1)`, ArgumentTypeString, args{-1, false}, "-1", false},
		{`string(15.2)`, ArgumentTypeString, args{15.2, false}, "15.2", false},
		{`string(-8)`, ArgumentTypeString, args{-8, false}, "-8", false},
		{`string("-5.8")`, ArgumentTypeString, args{"-5.8", false}, "-5.8", false},
		{`string("false")`, ArgumentTypeString, args{"false", false}, "false", false},
		{`string("true")`, ArgumentTypeString, args{"true", false}, "true", false},
		{`string("foo")`, ArgumentTypeString, args{"foo", false}, "foo", false},
		{`string("2s")`, ArgumentTypeString, args{"2s", false}, "2s", false},
		{`string("1.5h25s")`, ArgumentTypeString, args{"1.5h25s", false}, "1.5h25s", false},
		{`string("-15m8s")`, ArgumentTypeString, args{"-15m8s", false}, "-15m8s", false},

		{`incompatible_milti_boolean_type`, ArgumentTypeBoolean, args{nil, true}, nil, true},
		{`milti_boolean([]bool{false,true})`, ArgumentTypeBoolean, args{[]bool{false, true}, true}, []bool{false, true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := tt.t.Parse(tt.args.v, tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgumentType.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("ArgumentType.Parse() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestArgument_Match(t *testing.T) {
	type fields struct {
		Name        string
		Type        ArgumentType
		Description string
		Required    bool
		RegExp      string
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", fields{RegExp: "^[1-9][0-9]*(\\.[0-9]*[1-9])?$"}, args{float64(1.0)}, true, false},
		{"", fields{RegExp: "^[1-9][0-9]*(\\.[0-9]*[1-9])?$"}, args{float64(1.01)}, true, false},
		{"", fields{RegExp: "^[1-9][0-9]*(\\.[0-9]*[1-9])?$"}, args{float64(10.01)}, true, false},
		{"", fields{RegExp: "^[1-9][0-9]*(\\.[0-9]*[1-9])?$"}, args{"10.01"}, true, false},
		{"", fields{RegExp: "^[1-9][0-9]*(\\.[0-9]*[1-9])?$"}, args{"1.0"}, false, false},
		{"", fields{RegExp: "^([1-9][0-9]*(\\.[0-9]*[1-9])?$"}, args{"1.0"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := Argument{
				Name:        tt.fields.Name,
				Type:        tt.fields.Type,
				Description: tt.fields.Description,
				Required:    tt.fields.Required,
				RegExp:      tt.fields.RegExp,
			}
			got, err := arg.Match(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Argument.Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Argument.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
