package sedoc

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/stdatiks/go-sedoc/types"
)

// ArgumentParser is parser for Argument
type ArgumentParser func(v interface{}, list bool) (r interface{}, err error)

var argumentParsers = map[ArgumentType]ArgumentParser{
	ArgumentTypeBoolean:  parseBoolean,
	ArgumentTypeDuration: parseDuration,
	ArgumentTypeUUID:     parseUUID,
	ArgumentTypeFloat:    parseFloat,
	ArgumentTypeInteger:  parseInteger,
	ArgumentTypeString:   parseString,
	ArgumentTypeTime:     parseTime,
}

// SetArgumentParser func
func SetArgumentParser(t ArgumentType, p ArgumentParser) {
	argumentParsers[t] = p
}

// Parse func
func (t ArgumentType) Parse(v interface{}, list bool) (r interface{}, err error) {
	if p, ok := argumentParsers[t]; !ok {
		err = fmt.Errorf("incompatible type: %s (%v)", t, v)
	} else if p == nil {
		err = fmt.Errorf("nil parser: %s (%v)", t, v)
	} else {
		r, err = p(v, list)
		if err != nil {
			err = fmt.Errorf("api: %v", err)
		}
	}
	return
}

func parseBoolean(v interface{}, list bool) (r interface{}, err error) {
	if list {
		if v == nil {
			err = fmt.Errorf("list can not be nil")
			return
		}
		switch v := v.(type) {
		case []bool:
			r = v
		case []interface{}:
			result := make([]bool, len(v))
			for idx := range v {
				var item interface{}
				if item, err = parseBoolean(v[idx], false); err != nil {
					return
				}
				result[idx] = item.(bool)
			}
			r = result
		default:
			err = fmt.Errorf("incompatible boolean list type: %T (%v)", v, v)
		}
		return
	}
	switch v := v.(type) {
	case bool:
		r = v
	case int:
		r = (v != 0)
	case float64:
		r = (v != 0)
	case string:
		r = (v == "true")
	default:
		err = fmt.Errorf("incompatible boolean type: %T (%v)", v, v)
		r = false
	}
	return
}
func parseDuration(v interface{}, list bool) (r interface{}, err error) {
	if list {
		if v == nil {
			err = fmt.Errorf("list can not be nil")
			return
		}
		switch v := v.(type) {
		case []types.Duration:
			r = v
		case []interface{}:
			result := make([]types.Duration, len(v))
			for idx := range v {
				var item interface{}
				if item, err = parseDuration(v[idx], false); err != nil {
					return
				}
				result[idx] = item.(types.Duration)
			}
			r = result
		default:
			err = fmt.Errorf("incompatible duration list type: %T (%v)", v, v)
		}
		return
	}
	switch v := v.(type) {
	case types.Duration:
		r = v
	case time.Duration:
		r = types.Duration(v)
	case int:
		r = types.Duration(v)
	case string:
		var t time.Duration
		t, err = time.ParseDuration(v)
		r = types.Duration(t)
	default:
		err = fmt.Errorf("incompatible duration type: %T (%v)", v, v)
		r = types.Duration(0)
	}
	return
}
func parseTime(v interface{}, list bool) (r interface{}, err error) {
	if list {
		if v == nil {
			err = fmt.Errorf("list can not be nil")
			return
		}
		switch v := v.(type) {
		case []time.Time:
			r = v
		case []interface{}:
			result := make([]time.Time, len(v))
			for idx := range v {
				var item interface{}
				if item, err = parseTime(v[idx], false); err != nil {
					return
				}
				result[idx] = item.(time.Time)
			}
			r = result
		default:
			err = fmt.Errorf("incompatible datetime list type: %T (%v)", v, v)
		}
		return
	}
	switch v := v.(type) {
	case time.Time:
		r = v
	case string:
		r, err = time.Parse(time.RFC3339, v)
	default:
		err = fmt.Errorf("incompatible datetime type: %T (%v)", v, v)
		r = time.Time{}
	}
	return
}
func parseUUID(v interface{}, list bool) (r interface{}, err error) {
	if list {
		if v == nil {
			err = fmt.Errorf("list can not be nil")
			return
		}
		switch v := v.(type) {
		case []types.UUID:
			r = v
		case []interface{}:
			result := make([]types.UUID, len(v))
			for idx := range v {
				var item interface{}
				if item, err = parseUUID(v[idx], false); err != nil {
					return
				}
				result[idx] = item.(types.UUID)
			}
			r = result
		default:
			err = fmt.Errorf("incompatible uuid list type: %T (%v)", v, v)
		}
		return
	}
	switch v := v.(type) {
	case types.UUID:
		r = v
	case uuid.UUID:
		u := types.UUID{}
		u.UUID = v
		r = u
	case string:
		u := types.UUID{}
		u.UUID, err = uuid.Parse(v)
		r = u
	default:
		err = fmt.Errorf("incompatible uuid type: %T (%v)", v, v)
		u := types.UUID{}
		u.UUID = uuid.Nil
		r = u
	}
	return
}

func parseInteger(v interface{}, list bool) (r interface{}, err error) {
	if list {
		if v == nil {
			err = fmt.Errorf("list can not be nil")
			return
		}
		switch v := v.(type) {
		case []int:
			r = v
		case []interface{}:
			result := make([]int, len(v))
			for idx := range v {
				var item interface{}
				if item, err = parseInteger(v[idx], false); err != nil {
					return
				}
				result[idx] = item.(int)
			}
			r = result
		default:
			err = fmt.Errorf("incompatible integer list type: %T (%v)", v, v)
		}
		return
	}
	switch v := v.(type) {
	case bool:
		r = 0
		if v {
			r = 1
		}
	case int:
		r = v
	case float64:
		r = int(v)
	case string:
		var i int64
		i, err = strconv.ParseInt(v, 10, 8)
		r = int(i)
	default:
		err = fmt.Errorf("incompatible integer type: %T (%v)", v, v)
		r = 0
	}
	return
}

// func parseIntegerDefault(v interface{}, def int) int {
//  if val, err := parseInteger(v); err != nil {
//    return val.(int)
//  }
//  return def
// }

func parseFloat(v interface{}, list bool) (r interface{}, err error) {
	if list {
		if v == nil {
			err = fmt.Errorf("list can not be nil")
			return
		}
		switch v := v.(type) {
		case []float64:
			r = v
		case []interface{}:
			result := make([]float64, len(v))
			for idx := range v {
				var item interface{}
				if item, err = parseFloat(v[idx], false); err != nil {
					return
				}
				result[idx] = item.(float64)
			}
			r = result
		default:
			err = fmt.Errorf("incompatible float list type: %T (%v)", v, v)
		}
		return
	}
	switch v := v.(type) {
	case bool:
		r = float64(0)
		if v {
			r = float64(1)
		}
	case int:
		r = float64(v)
	case float64:
		r = v
	case string:
		r, err = strconv.ParseFloat(v, 8)
	default:
		err = fmt.Errorf("incompatible float type: %T (%v)", v, v)
		r = float64(0)
	}
	return
}
func parseString(v interface{}, list bool) (r interface{}, err error) {
	if list {
		if v == nil {
			err = fmt.Errorf("list can not be nil")
			return
		}
		switch v := v.(type) {
		case []string:
			r = v
		case []interface{}:
			result := make([]string, len(v))
			for idx := range v {
				var item interface{}
				if item, err = parseString(v[idx], false); err != nil {
					return
				}
				result[idx] = item.(string)
			}
			r = result
		default:
			err = fmt.Errorf("incompatible string list type: %T (%v)", v, v)
		}
		return
	}
	switch v := v.(type) {
	case bool:
		r = fmt.Sprintf("%t", v)
	case int:
		if v != 0 {
			r = strconv.FormatInt(int64(v), 10)
		} else {
			r = "0"
		}
	case float64:
		if v != 0 {
			r = strconv.FormatFloat(v, 'f', -1, 64)
		} else {
			r = "0.0"
		}
	case string:
		r = v
	default:
		err = fmt.Errorf("incompatible string type: %T (%v)", v, v)
		r = ""
	}
	return
}
