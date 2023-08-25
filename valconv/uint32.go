package valconv

import (
	"reflect"
	"strconv"
)

type ValueUInt32 struct {
	Value uint32
	Err   error
}

func UInt32(v uint32) ValueUInt32 {
	return ValueUInt32{v, nil}
}

func AnyAsUInt32(v any) (uint32, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(uint32); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return uint32(val.Uint()), true
	}
	return 0, false
}

func AnyToUInt32(v any) (uint32, bool) {
	if v2, ok := AnyAsUInt32(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return uint32(v2), false
	}
	return uint32(v2), true
}

func (o ValueUInt32) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatUint(uint64(o.Value), 10)
	return String(v2)
}

func (o ValueUInt32) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func UInt32ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v uint32) T {
	return T(v)
}

func (o ValueUInt32) ToInt() ValueInt {
	return ValueInt{UInt32ToNumber[int](o.Value), o.Err}
}

func (o ValueUInt32) ToInt8() ValueInt8 {
	return ValueInt8{UInt32ToNumber[int8](o.Value), o.Err}
}

func (o ValueUInt32) ToInt16() ValueInt16 {
	return ValueInt16{UInt32ToNumber[int16](o.Value), o.Err}
}

func (o ValueUInt32) ToInt32() ValueInt32 {
	return ValueInt32{UInt32ToNumber[int32](o.Value), o.Err}
}

func (o ValueUInt32) ToInt64() ValueInt64 {
	return ValueInt64{UInt32ToNumber[int64](o.Value), o.Err}
}

func (o ValueUInt32) ToUInt() ValueUInt {
	return ValueUInt{UInt32ToNumber[uint](o.Value), o.Err}
}

func (o ValueUInt32) ToUInt8() ValueUInt8 {
	return ValueUInt8{UInt32ToNumber[uint8](o.Value), o.Err}
}

func (o ValueUInt32) ToUInt16() ValueUInt16 {
	return ValueUInt16{UInt32ToNumber[uint16](o.Value), o.Err}
}
func (o ValueUInt32) ToUInt32() ValueUInt32 {
	return ValueUInt32{UInt32ToNumber[uint32](o.Value), o.Err}
}
func (o ValueUInt32) ToUInt64() ValueUInt64 {
	return ValueUInt64{UInt32ToNumber[uint64](o.Value), o.Err}
}
func (o ValueUInt32) ToFloat32() ValueFloat32 {
	return ValueFloat32{UInt32ToNumber[float32](o.Value), o.Err}
}
func (o ValueUInt32) ToFloat64() ValueFloat64 {
	return ValueFloat64{UInt32ToNumber[float64](o.Value), o.Err}
}
