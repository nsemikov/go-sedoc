package sedoc

import (
	"net/url"
	"strings"
)

// RequestFromURL func
func (api API) RequestFromURL(r *Request, u *url.URL) *Request {
	request := r
	if request == nil {
		request = NewRequest()
	}
	api.addToRequest(u.Query(), request)
	return request
}

func (api API) addToRequest(values url.Values, request *Request) {
	cmd := api.GetCommand(request.Command)
	if len(values) > 0 {
		// get request params from uri
		for key, vals := range values {
			if len(vals) == 0 {
				continue
			}
			if strings.Index(key, api.PrefixSet) == 0 {
				name := key[len(api.PrefixSet):]
				val := interface{}(vals[0])
				if multiple(cmd.Set, name) {
					arr := make([]interface{}, len(vals))
					for idx := range vals {
						arr[idx] = vals[idx]
					}
					val = arr
				}
				addToRequestSet(&request.Set, name, val)
			} else if strings.Index(key, api.PrefixWhere) == 0 {
				name := key[len(api.PrefixWhere):]
				val := interface{}(vals[0])
				if multiple(cmd.Where, name) {
					arr := make([]interface{}, len(vals))
					for idx := range vals {
						arr[idx] = vals[idx]
					}
					val = arr
				}
				addToRequestWhere(&request.Where, 0, name, val)
			} else if strings.Index(key, api.PrefixArguments) == 0 {
				name := key[len(api.PrefixArguments):]
				val := interface{}(vals[0])
				if multiple(cmd.Arguments, name) {
					arr := make([]interface{}, len(vals))
					for idx := range vals {
						arr[idx] = vals[idx]
					}
					val = arr
				}
				addToRequestArguments(&request.Arguments, name, val)
			}
		}
	}
}

func multiple(args Arguments, name string) bool {
	arg, err := args.Get(name)
	if err != nil || !arg.Multiple {
		return false
	}
	return true
}

func addToRequestSet(set *InterfaceMap, name string, value interface{}) {
	if *set == nil {
		*set = make(InterfaceMap)
	}
	(*set)[name] = value
}

func addToRequestArguments(arguments *InterfaceMap, name string, value interface{}) {
	if *arguments == nil {
		*arguments = make(InterfaceMap)
	}
	(*arguments)[name] = value
}

func addToRequestWhere(where *[]InterfaceMap, idx int, name string, value interface{}) {
	if *where == nil || len(*where) == 0 {
		*where = make([]InterfaceMap, 1)
		(*where)[0] = make(InterfaceMap)
	}
	(*where)[idx][name] = value
}
