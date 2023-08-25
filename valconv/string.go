package valconv

import (
	"fmt"
	"strconv"
	"time"
)

type ValueString struct {
	Value string
	Err   error
}

func String(v string) ValueString {
	return ValueString{v, nil}
}

func AnyString(v any) ValueString {
	return String(AnyToString(v))
}

func AnyToString(v any) string {
	var str string
	if v != nil {
		switch value := v.(type) {
		case int:
			str = strconv.Itoa(value)
		case int8:
			str = strconv.Itoa(int(value))
		case int16:
			str = strconv.Itoa(int(value))
		case int32: // same as `rune`
			str = strconv.Itoa(int(value))
		case int64:
			str = strconv.FormatInt(value, 10)
		case uint:
			str = strconv.FormatUint(uint64(value), 10)
		case uint8:
			str = strconv.FormatUint(uint64(value), 10)
		case uint16:
			str = strconv.FormatUint(uint64(value), 10)
		case uint32:
			str = strconv.FormatUint(uint64(value), 10)
		case uint64:
			str = strconv.FormatUint(value, 10)
		case float32:
			str = strconv.FormatFloat(float64(value), 'f', -1, 32)
		case float64:
			str = strconv.FormatFloat(value, 'f', -1, 64)
		case bool:
			str = strconv.FormatBool(value)
		case string:
			str = value
		case []byte:
			str = string(value)
		case time.Duration:
			str = strconv.FormatInt(int64(value), 10)
		case fmt.Stringer:
			str = value.String()
		case error:
			str = value.Error()
		default:
			str = fmt.Sprint(v)
		}
	}
	return str
}
