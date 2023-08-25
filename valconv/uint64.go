package valconv

import (
	"reflect"
	"strconv"
)

type ValueUInt64 struct {
	Value uint64
	Err   error
}

func UInt64(v uint64) ValueUInt64 {
	return ValueUInt64{v, nil}
}

func AnyAsUInt64(v any) (uint64, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(uint64); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return uint64(val.Uint()), true
	}
	return 0, false
}

func AnyToUInt64(v any) (uint64, bool) {
	if v2, ok := AnyAsUInt64(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return uint64(v2), false
	}
	return uint64(v2), true
}

func (o ValueUInt64) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatUint(uint64(o.Value), 10)
	return String(v2)
}

func (o ValueUInt64) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func UInt64ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v uint64) T {
	return T(v)
}

func (o ValueUInt64) ToInt() ValueInt {
	return ValueInt{UInt64ToNumber[int](o.Value), o.Err}
}

func (o ValueUInt64) ToInt8() ValueInt8 {
	return ValueInt8{UInt64ToNumber[int8](o.Value), o.Err}
}

func (o ValueUInt64) ToInt16() ValueInt16 {
	return ValueInt16{UInt64ToNumber[int16](o.Value), o.Err}
}

func (o ValueUInt64) ToInt32() ValueInt32 {
	return ValueInt32{UInt64ToNumber[int32](o.Value), o.Err}
}

func (o ValueUInt64) ToInt64() ValueInt64 {
	return ValueInt64{UInt64ToNumber[int64](o.Value), o.Err}
}

func (o ValueUInt64) ToUInt() ValueUInt {
	return ValueUInt{UInt64ToNumber[uint](o.Value), o.Err}
}

func (o ValueUInt64) ToUInt168() ValueUInt8 {
	return ValueUInt8{UInt64ToNumber[uint8](o.Value), o.Err}
}

func (o ValueUInt64) ToUInt16() ValueUInt16 {
	return ValueUInt16{UInt64ToNumber[uint16](o.Value), o.Err}
}
func (o ValueUInt64) ToUInt32() ValueUInt32 {
	return ValueUInt32{UInt64ToNumber[uint32](o.Value), o.Err}
}
func (o ValueUInt64) ToUInt64() ValueUInt64 {
	return ValueUInt64{UInt64ToNumber[uint64](o.Value), o.Err}
}
func (o ValueUInt64) ToFloat32() ValueFloat32 {
	return ValueFloat32{UInt64ToNumber[float32](o.Value), o.Err}
}
func (o ValueUInt64) ToFloat64() ValueFloat64 {
	return ValueFloat64{UInt64ToNumber[float64](o.Value), o.Err}
}
