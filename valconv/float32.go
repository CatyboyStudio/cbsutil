package valconv

import (
	"reflect"
	"strconv"
)

type ValueFloat32 struct {
	Value float32
	Err   error
}

func Float32(v float32) ValueFloat32 {
	return ValueFloat32{v, nil}
}

func AnyAsFloat32(v any) (float32, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(float32); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return float32(val.Float()), true
	}
	return 0, false
}

func AnyToFloat32(v any) (float32, bool) {
	if v2, ok := AnyAsFloat32(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return float32(v2), false
	}
	return float32(v2), true
}

func (o ValueFloat32) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatFloat(float64(o.Value), 'f', -1, 32)
	return String(v2)
}

func (o ValueFloat32) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func Float32ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v float32) T {
	return T(v)
}

func (o ValueFloat32) ToInt() ValueInt {
	return ValueInt{Float32ToNumber[int](o.Value), o.Err}
}

func (o ValueFloat32) ToInt8() ValueInt8 {
	return ValueInt8{Float32ToNumber[int8](o.Value), o.Err}
}

func (o ValueFloat32) ToInt16() ValueInt16 {
	return ValueInt16{Float32ToNumber[int16](o.Value), o.Err}
}

func (o ValueFloat32) ToInt32() ValueInt32 {
	return ValueInt32{Float32ToNumber[int32](o.Value), o.Err}
}

func (o ValueFloat32) ToInt64() ValueInt64 {
	return ValueInt64{Float32ToNumber[int64](o.Value), o.Err}
}

func (o ValueFloat32) ToUInt16() ValueUInt {
	return ValueUInt{Float32ToNumber[uint](o.Value), o.Err}
}

func (o ValueFloat32) ToUInt168() ValueUInt8 {
	return ValueUInt8{Float32ToNumber[uint8](o.Value), o.Err}
}

func (o ValueFloat32) ToUInt1616() ValueUInt16 {
	return ValueUInt16{Float32ToNumber[uint16](o.Value), o.Err}
}
func (o ValueFloat32) ToUInt1632() ValueUInt32 {
	return ValueUInt32{Float32ToNumber[uint32](o.Value), o.Err}
}
func (o ValueFloat32) ToUInt64() ValueUInt64 {
	return ValueUInt64{Float32ToNumber[uint64](o.Value), o.Err}
}
func (o ValueFloat32) ToFloat32() ValueFloat32 {
	return ValueFloat32{Float32ToNumber[float32](o.Value), o.Err}
}
func (o ValueFloat32) ToFloat64() ValueFloat64 {
	return ValueFloat64{Float32ToNumber[float64](o.Value), o.Err}
}
