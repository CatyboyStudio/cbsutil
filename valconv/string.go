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

func (o ValueString) ToBool() ValueBool {
	if o.Err != nil {
		return ValueBool{false, o.Err}
	}
	v, err := strconv.ParseBool(o.Value)
	return ValueBool{v, err}
}

func (o ValueString) ToInt() ValueInt {
	if o.Err != nil {
		return ValueInt{0, o.Err}
	}
	v, err := strconv.ParseInt(o.Value, 10, 0)
	return ValueInt{int(v), err}
}

func (o ValueString) ToInt8() ValueInt8 {
	if o.Err != nil {
		return ValueInt8{0, o.Err}
	}
	v, err := strconv.ParseInt(o.Value, 10, 8)
	return ValueInt8{int8(v), err}
}

func (o ValueString) ToInt16() ValueInt16 {
	if o.Err != nil {
		return ValueInt16{0, o.Err}
	}
	v, err := strconv.ParseInt(o.Value, 10, 16)
	return ValueInt16{int16(v), err}
}

func (o ValueString) ToInt32() ValueInt32 {
	if o.Err != nil {
		return ValueInt32{0, o.Err}
	}
	v, err := strconv.ParseInt(o.Value, 10, 32)
	return ValueInt32{int32(v), err}
}

func (o ValueString) ToInt64() ValueInt64 {
	if o.Err != nil {
		return ValueInt64{0, o.Err}
	}
	v, err := strconv.ParseInt(o.Value, 10, 64)
	return ValueInt64{int64(v), err}
}

func (o ValueString) ToUInt() ValueUInt {
	if o.Err != nil {
		return ValueUInt{0, o.Err}
	}
	v, err := strconv.ParseUint(o.Value, 10, 0)
	return ValueUInt{uint(v), err}
}

func (o ValueString) ToUInt8() ValueUInt8 {
	if o.Err != nil {
		return ValueUInt8{0, o.Err}
	}
	v, err := strconv.ParseUint(o.Value, 10, 8)
	return ValueUInt8{uint8(v), err}
}

func (o ValueString) ToUInt16() ValueUInt16 {
	if o.Err != nil {
		return ValueUInt16{0, o.Err}
	}
	v, err := strconv.ParseUint(o.Value, 10, 16)
	return ValueUInt16{uint16(v), err}
}
func (o ValueString) ToUInt32() ValueUInt32 {
	if o.Err != nil {
		return ValueUInt32{0, o.Err}
	}
	v, err := strconv.ParseUint(o.Value, 10, 32)
	return ValueUInt32{uint32(v), err}
}
func (o ValueString) ToUInt64() ValueUInt64 {
	if o.Err != nil {
		return ValueUInt64{0, o.Err}
	}
	v, err := strconv.ParseUint(o.Value, 10, 64)
	return ValueUInt64{uint64(v), err}
}
func (o ValueString) ToFloat32() ValueFloat32 {
	if o.Err != nil {
		return ValueFloat32{0, o.Err}
	}
	v, err := strconv.ParseFloat(o.Value, 32)
	return ValueFloat32{float32(v), err}
}
func (o ValueString) ToFloat64() ValueFloat64 {
	if o.Err != nil {
		return ValueFloat64{0, o.Err}
	}
	v, err := strconv.ParseFloat(o.Value, 64)
	return ValueFloat64{float64(v), err}
}
