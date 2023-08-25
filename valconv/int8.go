package valconv

import (
	"reflect"
	"strconv"
)

type ValueInt8 struct {
	Value int8
	Err   error
}

func Int8(v int8) ValueInt8 {
	return ValueInt8{v, nil}
}

func AnyAsInt8(v any) (int8, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(int8); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return int8(val.Int()), true
	}
	return 0, false
}

func AnyToInt8(v any) (int8, bool) {
	if v2, ok := AnyAsInt8(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return int8(v2), false
	}
	return int8(v2), true
}

func (o ValueInt8) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatInt(int64(o.Value), 10)
	return String(v2)
}

func (o ValueInt8) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func Int8ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v int8) T {
	return T(v)
}

func (o ValueInt8) ToInt8() ValueInt {
	return ValueInt{Int8ToNumber[int](o.Value), o.Err}
}

func (o ValueInt8) ToInt88() ValueInt8 {
	return ValueInt8{Int8ToNumber[int8](o.Value), o.Err}
}

func (o ValueInt8) ToInt816() ValueInt16 {
	return ValueInt16{Int8ToNumber[int16](o.Value), o.Err}
}

func (o ValueInt8) ToInt832() ValueInt32 {
	return ValueInt32{Int8ToNumber[int32](o.Value), o.Err}
}

func (o ValueInt8) ToInt864() ValueInt64 {
	return ValueInt64{Int8ToNumber[int64](o.Value), o.Err}
}

func (o ValueInt8) ToUInt8() ValueUInt {
	return ValueUInt{Int8ToNumber[uint](o.Value), o.Err}
}

func (o ValueInt8) ToUInt88() ValueUInt8 {
	return ValueUInt8{Int8ToNumber[uint8](o.Value), o.Err}
}

func (o ValueInt8) ToUInt816() ValueUInt16 {
	return ValueUInt16{Int8ToNumber[uint16](o.Value), o.Err}
}
func (o ValueInt8) ToUInt832() ValueUInt32 {
	return ValueUInt32{Int8ToNumber[uint32](o.Value), o.Err}
}
func (o ValueInt8) ToUInt864() ValueUInt64 {
	return ValueUInt64{Int8ToNumber[uint64](o.Value), o.Err}
}
func (o ValueInt8) ToFloat32() ValueFloat32 {
	return ValueFloat32{Int8ToNumber[float32](o.Value), o.Err}
}
func (o ValueInt8) ToFloat64() ValueFloat64 {
	return ValueFloat64{Int8ToNumber[float64](o.Value), o.Err}
}
