package valconv

import (
	"reflect"
	"strconv"
)

type ValueInt struct {
	Value int
	Err   error
}

func Int(v int) ValueInt {
	return ValueInt{v, nil}
}

func AnyAsInt(v any) (int, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(int); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return int(val.Int()), true
	}
	return 0, false
}

func AnyToInt(v any) (int, bool) {
	if v2, ok := AnyAsInt(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.Atoi(s)
	if err != nil {
		return v2, false
	}
	return v2, true
}

func (o ValueInt) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.Itoa(o.Value)
	return String(v2)
}

func (o ValueInt) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func IntToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v int) T {
	return T(v)
}

func (o ValueInt) ToInt() ValueInt {
	return ValueInt{IntToNumber[int](o.Value), o.Err}
}

func (o ValueInt) ToInt8() ValueInt8 {
	return ValueInt8{IntToNumber[int8](o.Value), o.Err}
}

func (o ValueInt) ToInt16() ValueInt16 {
	return ValueInt16{IntToNumber[int16](o.Value), o.Err}
}

func (o ValueInt) ToInt32() ValueInt32 {
	return ValueInt32{IntToNumber[int32](o.Value), o.Err}
}

func (o ValueInt) ToInt64() ValueInt64 {
	return ValueInt64{IntToNumber[int64](o.Value), o.Err}
}

func (o ValueInt) ToUInt() ValueUInt {
	return ValueUInt{IntToNumber[uint](o.Value), o.Err}
}

func (o ValueInt) ToUInt8() ValueUInt8 {
	return ValueUInt8{IntToNumber[uint8](o.Value), o.Err}
}

func (o ValueInt) ToUInt16() ValueUInt16 {
	return ValueUInt16{IntToNumber[uint16](o.Value), o.Err}
}
func (o ValueInt) ToUInt32() ValueUInt32 {
	return ValueUInt32{IntToNumber[uint32](o.Value), o.Err}
}
func (o ValueInt) ToUInt64() ValueUInt64 {
	return ValueUInt64{IntToNumber[uint64](o.Value), o.Err}
}
func (o ValueInt) ToFloat32() ValueFloat32 {
	return ValueFloat32{IntToNumber[float32](o.Value), o.Err}
}
func (o ValueInt) ToFloat64() ValueFloat64 {
	return ValueFloat64{IntToNumber[float64](o.Value), o.Err}
}
