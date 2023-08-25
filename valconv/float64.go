package valconv

import (
	"reflect"
	"strconv"
)

type ValueFloat64 struct {
	Value float64
	Err   error
}

func Float64(v float64) ValueFloat64 {
	return ValueFloat64{v, nil}
}

func AnyAsFloat64(v any) (float64, bool) {
	if v == nil {
		return 0, true
	}
	if v2, ok := v.(float64); ok {
		return v2, true
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Int16,
		reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8,
		reflect.Float32, reflect.Float64:
		return float64(val.Float()), true
	}
	return 0, false
}

func AnyToFloat64(v any) (float64, bool) {
	if v2, ok := AnyAsFloat64(v); ok {
		return v2, ok
	}
	s := AnyToString(v)
	v2, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(v2), false
	}
	return float64(v2), true
}

func (o ValueFloat64) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	v2 := strconv.FormatFloat(float64(o.Value), 'f', -1, 64)
	return String(v2)
}

func (o ValueFloat64) ToBool() ValueBool {
	v := true
	if o.Value == 0 {
		v = false
	}
	return ValueBool{v, o.Err}
}

func Float64ToNumber[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](v float64) T {
	return T(v)
}

func (o ValueFloat64) ToInt() ValueInt {
	return ValueInt{Float64ToNumber[int](o.Value), o.Err}
}

func (o ValueFloat64) ToInt8() ValueInt8 {
	return ValueInt8{Float64ToNumber[int8](o.Value), o.Err}
}

func (o ValueFloat64) ToInt16() ValueInt16 {
	return ValueInt16{Float64ToNumber[int16](o.Value), o.Err}
}

func (o ValueFloat64) ToInt32() ValueInt32 {
	return ValueInt32{Float64ToNumber[int32](o.Value), o.Err}
}

func (o ValueFloat64) ToInt64() ValueInt64 {
	return ValueInt64{Float64ToNumber[int64](o.Value), o.Err}
}

func (o ValueFloat64) ToUInt16() ValueUInt {
	return ValueUInt{Float64ToNumber[uint](o.Value), o.Err}
}

func (o ValueFloat64) ToUInt168() ValueUInt8 {
	return ValueUInt8{Float64ToNumber[uint8](o.Value), o.Err}
}

func (o ValueFloat64) ToUInt1616() ValueUInt16 {
	return ValueUInt16{Float64ToNumber[uint16](o.Value), o.Err}
}
func (o ValueFloat64) ToUInt1632() ValueUInt32 {
	return ValueUInt32{Float64ToNumber[uint32](o.Value), o.Err}
}
func (o ValueFloat64) ToUInt64() ValueUInt64 {
	return ValueUInt64{Float64ToNumber[uint64](o.Value), o.Err}
}
func (o ValueFloat64) ToFloat32() ValueFloat32 {
	return ValueFloat32{Float64ToNumber[float32](o.Value), o.Err}
}
func (o ValueFloat64) ToFloat64() ValueFloat64 {
	return ValueFloat64{Float64ToNumber[float64](o.Value), o.Err}
}
