package sedoc

import (
	"encoding/xml"
	"fmt"
)

// Command is api command
type Command struct {
	XMLName     xml.Name    `json:"-" xml:"command" yaml:"-"`
	Name        string      `json:"name" xml:"name,attr" yaml:"name"`
	Description string      `json:"description" xml:"description,attr" yaml:"description"`
	Arguments   Arguments   `json:"args,omitempty" xml:"args,omitempty" yaml:"args,omitempty"`
	Where       Arguments   `json:"where,omitempty" xml:"where,omitempty" yaml:"where,omitempty"`
	Set         Arguments   `json:"set,omitempty" xml:"set,omitempty" yaml:"set,omitempty"`
	Handler     HandlerFunc `json:"-" xml:"-" yaml:"-"`
	Examples    Examples    `json:"examples,omitempty" xml:"examples,omitempty" yaml:"examples,omitempty"`
}

// Commands is array of Command
type Commands []Command

// Contains Command
func (arr *Commands) Contains(name string) bool {
	_, err := arr.Get(name)
	return err == nil
}

// Add is add command to array
func (arr *Commands) Add(cmd Command) error {
	if arr.Contains(cmd.Name) {
		panic("command already exist: " + cmd.Name)
	}
	*arr = append(*arr, cmd)
	return nil
}

// Get is get command from array
func (arr *Commands) Get(name string) (Command, error) {
	for _, cmd := range *arr {
		if cmd.Name == name {
			return cmd, nil
		}
	}
	return Command{}, DefaultErrors.Get(ErrUnknownCommand)
}

// Remove is remove command from array
func (arr *Commands) Remove(name string) error {
	for idx, cmd := range *arr {
		if cmd.Name == name {
			*arr = append(append(Commands{}, (*arr)[:idx]...), (*arr)[idx+1:]...)
			return nil
		}
	}
	return DefaultErrors.Get(ErrUnknownCommand)
}

// MarshalXML for marshal into XML
func (arr Commands) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Space: "", Local: "commands"}
	a := []interface{}{}
	for _, item := range arr {
		a = append(a, item)
	}
	return MarshallerXML(a)(e, start)
}

// Valid func
func (cmd *Command) Valid() bool {
	return len(cmd.Name) > 0 && cmd.Handler != nil
}

func (cmd *Command) checkArguments(request *Request) (err error) {
	if err = checkArguments(&request.Arguments, cmd.Arguments, "args: "); err != nil {
		return
	}
	if err = checkArguments(&request.Set, cmd.Set, "set: "); err != nil {
		return
	}
	for idx, where := range request.Where {
		if err = checkArguments(&where, cmd.Where, fmt.Sprintf("where[%d]: ", idx)); err != nil {
			return
		}
		request.Where[idx] = where
	}
	return
}

func checkArguments(params *InterfaceMap, args Arguments, errPrefix string) (e error) {
	var err = &Error{}
	for _, arg := range args {
		if arg.Disabled {
			continue
		}
		_, ok := (*params)[arg.Name]
		if arg.Required && !ok {
			*err = DefaultErrors.Get(ErrRequiredArgumentMissing)
			err.Description = fmt.Sprintf("%s%s (%s)", errPrefix, err.Description, arg.Name)
			return err
		}
	}
	for name, val := range *params {
		arg, gerr := args.Get(name)
		if gerr != nil {
			*err = DefaultErrors.Get(ErrUnknownArgument)
			err.Description = fmt.Sprintf("%s%s (%s)", errPrefix, err.Description, name)
			return err
		}
		if val == nil {
			if !arg.Nullable {
				*err = DefaultErrors.Get(ErrInvalidArgumentValue)
				err.Description = fmt.Sprintf("%s%s (%s): %v", errPrefix, err.Description, name, "null")
				return err
			}
		} else if val, e = arg.Type.Parse(val, arg.Multiple); e != nil {
			*err = DefaultErrors.Get(ErrInvalidArgumentValue)
			err.Description = fmt.Sprintf("%s%s (%s): %v", errPrefix, err.Description, name, val)
			err.Internal = e
			return err
		}
		if len(arg.RegExp) > 0 {
			ok, e := arg.Match(val)
			if e != nil {
				*err = DefaultErrors.Get(ErrInvalidArgumentRegExp)
				err.Description = fmt.Sprintf("%s%s (%s)", errPrefix, err.Description, name)
				return err
			}
			if !ok {
				*err = DefaultErrors.Get(ErrArgumentRegExpMatchFails)
				err.Description = fmt.Sprintf("%s%s (%s, regexp: %s)", errPrefix, err.Description, name, arg.RegExp)
				return err
			}
		}
		(*params)[name] = val
	}
	return nil
}
