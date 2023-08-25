package valconv

import (
	"reflect"
	"strconv"
)

type ValueUInt16 struct {
	Value uint16
	Err   error
}

func UInt16(v uint16) ValueUInt16 {
	return ValueUInt16{v, nil}
}

func AnyAsUInt16(v any) (uint16, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(uint16); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return uint16(val.Uint()), true
	}
	return 0, false
}

func AnyToUInt16(v any) (uint16, bool) {
	if v2, ok := AnyAsUInt16(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return uint16(v2), false
	}
	return uint16(v2), true
}

func (o ValueUInt16) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatUint(uint64(o.Value), 10)
	return String(v2)
}

func (o ValueUInt16) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func UInt16ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v uint16) T {
	return T(v)
}

func (o ValueUInt16) ToInt() ValueInt {
	return ValueInt{UInt16ToNumber[int](o.Value), o.Err}
}

func (o ValueUInt16) ToInt8() ValueInt8 {
	return ValueInt8{UInt16ToNumber[int8](o.Value), o.Err}
}

func (o ValueUInt16) ToInt16() ValueInt16 {
	return ValueInt16{UInt16ToNumber[int16](o.Value), o.Err}
}

func (o ValueUInt16) ToInt32() ValueInt32 {
	return ValueInt32{UInt16ToNumber[int32](o.Value), o.Err}
}

func (o ValueUInt16) ToInt64() ValueInt64 {
	return ValueInt64{UInt16ToNumber[int64](o.Value), o.Err}
}

func (o ValueUInt16) ToUInt() ValueUInt {
	return ValueUInt{UInt16ToNumber[uint](o.Value), o.Err}
}

func (o ValueUInt16) ToUInt8() ValueUInt8 {
	return ValueUInt8{UInt16ToNumber[uint8](o.Value), o.Err}
}

func (o ValueUInt16) ToUInt16() ValueUInt16 {
	return ValueUInt16{UInt16ToNumber[uint16](o.Value), o.Err}
}
func (o ValueUInt16) ToUInt32() ValueUInt32 {
	return ValueUInt32{UInt16ToNumber[uint32](o.Value), o.Err}
}
func (o ValueUInt16) ToUInt64() ValueUInt64 {
	return ValueUInt64{UInt16ToNumber[uint64](o.Value), o.Err}
}
func (o ValueUInt16) ToFloat32() ValueFloat32 {
	return ValueFloat32{UInt16ToNumber[float32](o.Value), o.Err}
}
func (o ValueUInt16) ToFloat64() ValueFloat64 {
	return ValueFloat64{UInt16ToNumber[float64](o.Value), o.Err}
}
