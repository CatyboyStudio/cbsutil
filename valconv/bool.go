package valconv

import (
	"reflect"
	"strconv"
)

type ValueBool struct {
	Value bool
	Err   error
}

func Bool(v bool) ValueBool {
	return ValueBool{v, nil}
}

func AnyAsBool(v any) (bool, bool) {
	if v == nil {
		return false, true
	}
	if v2, ok := v.(bool); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Bool {
		return val.Bool(), true
	}
	return false, false
}

func AnyToBool(v any) (bool, bool) {
	if v2, ok := AnyAsBool(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseBool(s)
	if err != nil {
		return false, false
	}
	return v2, true
}

func (o ValueBool) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatBool(o.Value)
	return String(v2)
}

func (o ValueBool) ToBool() ValueBool {
	return o
}

func BoolToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](b bool) T {
	v := T(0)
	if b {
		v = 1
	}
	return v
}

func (o ValueBool) ToInt() ValueInt {
	return ValueInt{BoolToNumber[int](o.Value), o.Err}
}

func (o ValueBool) ToInt8() ValueInt8 {
	return ValueInt8{BoolToNumber[int8](o.Value), o.Err}
}

func (o ValueBool) ToInt16() ValueInt16 {
	return ValueInt16{BoolToNumber[int16](o.Value), o.Err}
}

func (o ValueBool) ToInt32() ValueInt32 {
	return ValueInt32{BoolToNumber[int32](o.Value), o.Err}
}

func (o ValueBool) ToInt64() ValueInt64 {
	return ValueInt64{BoolToNumber[int64](o.Value), o.Err}
}

func (o ValueBool) ToUInt() ValueUInt {
	return ValueUInt{BoolToNumber[uint](o.Value), o.Err}
}

func (o ValueBool) ToUInt8() ValueUInt8 {
	return ValueUInt8{BoolToNumber[uint8](o.Value), o.Err}
}

func (o ValueBool) ToUInt16() ValueUInt16 {
	return ValueUInt16{BoolToNumber[uint16](o.Value), o.Err}
}
func (o ValueBool) ToUInt32() ValueUInt32 {
	return ValueUInt32{BoolToNumber[uint32](o.Value), o.Err}
}
func (o ValueBool) ToUInt64() ValueUInt64 {
	return ValueUInt64{BoolToNumber[uint64](o.Value), o.Err}
}
func (o ValueBool) ToFloat32() ValueFloat32 {
	return ValueFloat32{BoolToNumber[float32](o.Value), o.Err}
}
func (o ValueBool) ToFloat64() ValueFloat64 {
	return ValueFloat64{BoolToNumber[float64](o.Value), o.Err}
}
