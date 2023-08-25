package valconv

import (
	"reflect"
	"strconv"
)

type ValueInt32 struct {
	Value int32
	Err   error
}

func Int32(v int32) ValueInt32 {
	return ValueInt32{v, nil}
}

func AnyAsInt32(v any) (int32, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(int32); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return int32(val.Int()), true
	}
	return 0, false
}

func AnyToInt32(v any) (int32, bool) {
	if v2, ok := AnyAsInt32(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return int32(v2), false
	}
	return int32(v2), true
}

func (o ValueInt32) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatInt(int64(o.Value), 10)
	return String(v2)
}

func (o ValueInt32) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func Int32ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v int32) T {
	return T(v)
}

func (o ValueInt32) ToInt() ValueInt {
	return ValueInt{Int32ToNumber[int](o.Value), o.Err}
}

func (o ValueInt32) ToInt8() ValueInt8 {
	return ValueInt8{Int32ToNumber[int8](o.Value), o.Err}
}

func (o ValueInt32) ToInt16() ValueInt16 {
	return ValueInt16{Int32ToNumber[int16](o.Value), o.Err}
}

func (o ValueInt32) ToInt32() ValueInt32 {
	return ValueInt32{Int32ToNumber[int32](o.Value), o.Err}
}

func (o ValueInt32) ToInt64() ValueInt64 {
	return ValueInt64{Int32ToNumber[int64](o.Value), o.Err}
}

func (o ValueInt32) ToUInt16() ValueUInt {
	return ValueUInt{Int32ToNumber[uint](o.Value), o.Err}
}

func (o ValueInt32) ToUInt168() ValueUInt8 {
	return ValueUInt8{Int32ToNumber[uint8](o.Value), o.Err}
}

func (o ValueInt32) ToUInt1616() ValueUInt16 {
	return ValueUInt16{Int32ToNumber[uint16](o.Value), o.Err}
}
func (o ValueInt32) ToUInt1632() ValueUInt32 {
	return ValueUInt32{Int32ToNumber[uint32](o.Value), o.Err}
}
func (o ValueInt32) ToUInt1664() ValueUInt64 {
	return ValueUInt64{Int32ToNumber[uint64](o.Value), o.Err}
}
func (o ValueInt32) ToFloat32() ValueFloat32 {
	return ValueFloat32{Int32ToNumber[float32](o.Value), o.Err}
}
func (o ValueInt32) ToFloat64() ValueFloat64 {
	return ValueFloat64{Int32ToNumber[float64](o.Value), o.Err}
}
