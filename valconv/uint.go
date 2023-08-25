package valconv

import (
	"reflect"
	"strconv"
)

type ValueUInt struct {
	Value uint
	Err   error
}

func UInt(v uint) ValueUInt {
	return ValueUInt{v, nil}
}

func AnyAsUInt(v any) (uint, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(uint); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return uint(val.Uint()), true
	}
	return 0, false
}

func AnyToUInt(v any) (uint, bool) {
	if v2, ok := AnyAsUInt(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return uint(v2), false
	}
	return uint(v2), true
}

func (o ValueUInt) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatUint(uint64(o.Value), 10)
	return String(v2)
}

func (o ValueUInt) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func UIntToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v uint) T {
	return T(v)
}

func (o ValueUInt) ToInt() ValueInt {
	return ValueInt{UIntToNumber[int](o.Value), o.Err}
}

func (o ValueUInt) ToInt8() ValueInt8 {
	return ValueInt8{UIntToNumber[int8](o.Value), o.Err}
}

func (o ValueUInt) ToInt16() ValueInt16 {
	return ValueInt16{UIntToNumber[int16](o.Value), o.Err}
}

func (o ValueUInt) ToInt32() ValueInt32 {
	return ValueInt32{UIntToNumber[int32](o.Value), o.Err}
}

func (o ValueUInt) ToInt64() ValueInt64 {
	return ValueInt64{UIntToNumber[int64](o.Value), o.Err}
}

func (o ValueUInt) ToUInt16() ValueUInt {
	return ValueUInt{UIntToNumber[uint](o.Value), o.Err}
}

func (o ValueUInt) ToUInt168() ValueUInt8 {
	return ValueUInt8{UIntToNumber[uint8](o.Value), o.Err}
}

func (o ValueUInt) ToUInt1616() ValueUInt16 {
	return ValueUInt16{UIntToNumber[uint16](o.Value), o.Err}
}
func (o ValueUInt) ToUInt1632() ValueUInt32 {
	return ValueUInt32{UIntToNumber[uint32](o.Value), o.Err}
}
func (o ValueUInt) ToUInt1664() ValueUInt64 {
	return ValueUInt64{UIntToNumber[uint64](o.Value), o.Err}
}
func (o ValueUInt) ToFloat32() ValueFloat32 {
	return ValueFloat32{UIntToNumber[float32](o.Value), o.Err}
}
func (o ValueUInt) ToFloat64() ValueFloat64 {
	return ValueFloat64{UIntToNumber[float64](o.Value), o.Err}
}
