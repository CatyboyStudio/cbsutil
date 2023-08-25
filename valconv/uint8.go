package valconv

import (
	"reflect"
	"strconv"
)

type ValueUInt8 struct {
	Value uint8
	Err   error
}

func UInt8(v uint8) ValueUInt8 {
	return ValueUInt8{v, nil}
}

func AnyAsUInt8(v any) (uint8, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(uint8); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return uint8(val.Uint()), true
	}
	return 0, false
}

func AnyToUInt8(v any) (uint8, bool) {
	if v2, ok := AnyAsUInt8(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return uint8(v2), false
	}
	return uint8(v2), true
}

func (o ValueUInt8) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatUint(uint64(o.Value), 10)
	return String(v2)
}

func (o ValueUInt8) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func UInt8ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v uint8) T {
	return T(v)
}

func (o ValueUInt8) ToInt() ValueInt {
	return ValueInt{UInt8ToNumber[int](o.Value), o.Err}
}

func (o ValueUInt8) ToInt8() ValueInt8 {
	return ValueInt8{UInt8ToNumber[int8](o.Value), o.Err}
}

func (o ValueUInt8) ToInt16() ValueInt16 {
	return ValueInt16{UInt8ToNumber[int16](o.Value), o.Err}
}

func (o ValueUInt8) ToInt32() ValueInt32 {
	return ValueInt32{UInt8ToNumber[int32](o.Value), o.Err}
}

func (o ValueUInt8) ToInt64() ValueInt64 {
	return ValueInt64{UInt8ToNumber[int64](o.Value), o.Err}
}

func (o ValueUInt8) ToUInt() ValueUInt {
	return ValueUInt{UInt8ToNumber[uint](o.Value), o.Err}
}

func (o ValueUInt8) ToUInt8() ValueUInt8 {
	return ValueUInt8{UInt8ToNumber[uint8](o.Value), o.Err}
}

func (o ValueUInt8) ToUInt16() ValueUInt16 {
	return ValueUInt16{UInt8ToNumber[uint16](o.Value), o.Err}
}
func (o ValueUInt8) ToUInt32() ValueUInt32 {
	return ValueUInt32{UInt8ToNumber[uint32](o.Value), o.Err}
}
func (o ValueUInt8) ToUInt64() ValueUInt64 {
	return ValueUInt64{UInt8ToNumber[uint64](o.Value), o.Err}
}
func (o ValueUInt8) ToFloat32() ValueFloat32 {
	return ValueFloat32{UInt8ToNumber[float32](o.Value), o.Err}
}
func (o ValueUInt8) ToFloat64() ValueFloat64 {
	return ValueFloat64{UInt8ToNumber[float64](o.Value), o.Err}
}
