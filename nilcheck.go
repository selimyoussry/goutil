package goutil

import "github.com/hippoai/goerr"

const (
	ErrMsgUnknownType = "ERR_UNKNOWN_TYPE"
)

func ErrUnknownType(v interface{}) error {
	return goerr.New(ErrMsgUnknownType, map[string]interface{}{
		"value": v,
	})
}

// NilCheck checks if a given value is nil
// it works only on basic types, not on slices or more complex data structures
func NilCheck(v interface{}) (bool, error) {

	switch v.(type) {

	case *string:
		val := v.(*string)
		return val == nil, nil

	case *bool:
		val := v.(*bool)
		return val == nil, nil

	case *float32:
		val := v.(*float32)
		return val == nil, nil

	case *float64:
		val := v.(*float64)
		return val == nil, nil

	case *int:
		val := v.(*int)
		return val == nil, nil

	case *int8:
		val := v.(*int8)
		return val == nil, nil

	case *int16:
		val := v.(*int16)
		return val == nil, nil

	case *int32:
		val := v.(*int32)
		return val == nil, nil

	case *int64:
		val := v.(*int64)
		return val == nil, nil

	case nil:
		return true, nil

	case string:
		return false, nil

	case bool:
		return false, nil

	case float32:
		return false, nil

	case float64:
		return false, nil

	case int:
		return false, nil

	case int8:
		return false, nil

	case int16:
		return false, nil

	case int32:
		return false, nil

	case int64:
		return false, nil

	default:
		return false, ErrUnknownType(v)
	}

}

// DefaultIfNull returns the default value if we get a nil value
func DefaultIfNull(v interface{}, defaultValue interface{}) interface{} {
	isNil, err := NilCheck(v)
	if (err != nil) || (isNil) {
		return defaultValue
	}

	return v
}
