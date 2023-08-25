package valconv

import (
	"reflect"
	"strconv"
)

type ValueInt16 struct {
	Value int16
	Err   error
}

func Int16(v int16) ValueInt16 {
	return ValueInt16{v, nil}
}

func AnyAsInt16(v any) (int16, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(int16); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return int16(val.Int()), true
	}
	return 0, false
}

func AnyToInt16(v any) (int16, bool) {
	if v2, ok := AnyAsInt16(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return int16(v2), false
	}
	return int16(v2), true
}

func (o ValueInt16) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatInt(int64(o.Value), 10)
	return String(v2)
}

func (o ValueInt16) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func Int16ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v int16) T {
	return T(v)
}

func (o ValueInt16) ToInt() ValueInt {
	return ValueInt{Int16ToNumber[int](o.Value), o.Err}
}

func (o ValueInt16) ToInt8() ValueInt8 {
	return ValueInt8{Int16ToNumber[int8](o.Value), o.Err}
}

func (o ValueInt16) ToInt16() ValueInt16 {
	return ValueInt16{Int16ToNumber[int16](o.Value), o.Err}
}

func (o ValueInt16) ToInt32() ValueInt32 {
	return ValueInt32{Int16ToNumber[int32](o.Value), o.Err}
}

func (o ValueInt16) ToInt64() ValueInt64 {
	return ValueInt64{Int16ToNumber[int64](o.Value), o.Err}
}

func (o ValueInt16) ToUInt16() ValueUInt {
	return ValueUInt{Int16ToNumber[uint](o.Value), o.Err}
}

func (o ValueInt16) ToUInt168() ValueUInt8 {
	return ValueUInt8{Int16ToNumber[uint8](o.Value), o.Err}
}

func (o ValueInt16) ToUInt1616() ValueUInt16 {
	return ValueUInt16{Int16ToNumber[uint16](o.Value), o.Err}
}
func (o ValueInt16) ToUInt1632() ValueUInt32 {
	return ValueUInt32{Int16ToNumber[uint32](o.Value), o.Err}
}
func (o ValueInt16) ToUInt1664() ValueUInt64 {
	return ValueUInt64{Int16ToNumber[uint64](o.Value), o.Err}
}
func (o ValueInt16) ToFloat32() ValueFloat32 {
	return ValueFloat32{Int16ToNumber[float32](o.Value), o.Err}
}
func (o ValueInt16) ToFloat64() ValueFloat64 {
	return ValueFloat64{Int16ToNumber[float64](o.Value), o.Err}
}
