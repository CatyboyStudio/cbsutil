package valconv

type ValueAny struct {
	Value any
	Err   error
}

func Any(v any) ValueAny {
	return ValueAny{v, nil}
}

func (o ValueAny) Bool() ValueBool {
	if o.Err != nil {
		return ValueBool{false, o.Err}
	}
	if v, ok := o.Value.(bool); ok {
		return Bool(v)
	}
	return ValueBool{false, ErrConvertType(o.Value, "bool")}
}

func (o ValueAny) AsBool() ValueBool {
	if o.Err != nil {
		return ValueBool{false, o.Err}
	}
	if v, ok := AnyAsBool(o.Value); ok {
		return Bool(v)
	}
	return ValueBool{false, ErrConvertType(o.Value, "bool")}
}

func (o ValueAny) ToBool() ValueBool {
	if o.Err != nil {
		return ValueBool{false, o.Err}
	}
	if v, ok := AnyAsBool(o.Value); ok {
		return Bool(v)
	}
	return ValueBool{false, ErrConvertType(o.Value, "bool")}
}

func (o ValueAny) String() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	if v, ok := o.Value.(string); ok {
		return String(v)
	}
	return ValueString{"", ErrConvertType(o.Value, "string")}
}

func (o ValueAny) ToString() ValueString {
	if o.Err != nil {
		return ValueString{"", o.Err}
	}
	return AnyString(o.Value)
}

func (o ValueAny) Int() ValueInt {
	if o.Err != nil {
		return ValueInt{0, o.Err}
	}
	if v, ok := o.Value.(int); ok {
		return Int(v)
	}
	return ValueInt{0, ErrConvertType(o.Value, "int")}
}

func (o ValueAny) AsInt() ValueInt {
	if o.Err != nil {
		return ValueInt{0, o.Err}
	}
	if v, ok := AnyAsInt(o.Value); ok {
		return Int(v)
	}
	return ValueInt{0, ErrConvertType(o.Value, "int")}
}

func (o ValueAny) ToInt() ValueInt {
	if o.Err != nil {
		return ValueInt{0, o.Err}
	}
	if v, ok := AnyAsInt(o.Value); ok {
		return Int(v)
	}
	return ValueInt{0, ErrConvertType(o.Value, "int")}
}

func (o ValueAny) Int8() ValueInt8 {
	if o.Err != nil {
		return ValueInt8{0, o.Err}
	}
	if v, ok := o.Value.(int8); ok {
		return Int8(v)
	}
	return ValueInt8{0, ErrConvertType(o.Value, "int8")}
}

func (o ValueAny) AsInt8() ValueInt8 {
	if o.Err != nil {
		return ValueInt8{0, o.Err}
	}
	if v, ok := AnyAsInt8(o.Value); ok {
		return Int8(v)
	}
	return ValueInt8{0, ErrConvertType(o.Value, "int8")}
}

func (o ValueAny) ToInt8() ValueInt8 {
	if o.Err != nil {
		return ValueInt8{0, o.Err}
	}
	if v, ok := AnyAsInt8(o.Value); ok {
		return Int8(v)
	}
	return ValueInt8{0, ErrConvertType(o.Value, "int8")}
}

func (o ValueAny) Int16() ValueInt16 {
	if o.Err != nil {
		return ValueInt16{0, o.Err}
	}
	if v, ok := o.Value.(int16); ok {
		return Int16(v)
	}
	return ValueInt16{0, ErrConvertType(o.Value, "int16")}
}

func (o ValueAny) AsInt16() ValueInt16 {
	if o.Err != nil {
		return ValueInt16{0, o.Err}
	}
	if v, ok := AnyAsInt16(o.Value); ok {
		return Int16(v)
	}
	return ValueInt16{0, ErrConvertType(o.Value, "int16")}
}

func (o ValueAny) ToInt16() ValueInt16 {
	if o.Err != nil {
		return ValueInt16{0, o.Err}
	}
	if v, ok := AnyAsInt16(o.Value); ok {
		return Int16(v)
	}
	return ValueInt16{0, ErrConvertType(o.Value, "int16")}
}

func (o ValueAny) Int32() ValueInt32 {
	if o.Err != nil {
		return ValueInt32{0, o.Err}
	}
	if v, ok := o.Value.(int32); ok {
		return Int32(v)
	}
	return ValueInt32{0, ErrConvertType(o.Value, "Int32")}
}

func (o ValueAny) AsInt32() ValueInt32 {
	if o.Err != nil {
		return ValueInt32{0, o.Err}
	}
	if v, ok := AnyAsInt32(o.Value); ok {
		return Int32(v)
	}
	return ValueInt32{0, ErrConvertType(o.Value, "Int32")}
}

func (o ValueAny) ToInt32() ValueInt32 {
	if o.Err != nil {
		return ValueInt32{0, o.Err}
	}
	if v, ok := AnyAsInt32(o.Value); ok {
		return Int32(v)
	}
	return ValueInt32{0, ErrConvertType(o.Value, "Int32")}
}

func (o ValueAny) Int64() ValueInt64 {
	if o.Err != nil {
		return ValueInt64{0, o.Err}
	}
	if v, ok := o.Value.(int64); ok {
		return Int64(v)
	}
	return ValueInt64{0, ErrConvertType(o.Value, "int64")}
}

func (o ValueAny) AsInt64() ValueInt64 {
	if o.Err != nil {
		return ValueInt64{0, o.Err}
	}
	if v, ok := AnyAsInt64(o.Value); ok {
		return Int64(v)
	}
	return ValueInt64{0, ErrConvertType(o.Value, "int64")}
}

func (o ValueAny) ToInt64() ValueInt64 {
	if o.Err != nil {
		return ValueInt64{0, o.Err}
	}
	if v, ok := AnyAsInt64(o.Value); ok {
		return Int64(v)
	}
	return ValueInt64{0, ErrConvertType(o.Value, "int64")}
}

func (o ValueAny) UInt() ValueUInt {
	if o.Err != nil {
		return ValueUInt{0, o.Err}
	}
	if v, ok := o.Value.(uint); ok {
		return UInt(v)
	}
	return ValueUInt{0, ErrConvertType(o.Value, "uint")}
}

func (o ValueAny) AsUInt() ValueUInt {
	if o.Err != nil {
		return ValueUInt{0, o.Err}
	}
	if v, ok := AnyAsUInt(o.Value); ok {
		return UInt(v)
	}
	return ValueUInt{0, ErrConvertType(o.Value, "uint")}
}

func (o ValueAny) ToUInt() ValueUInt {
	if o.Err != nil {
		return ValueUInt{0, o.Err}
	}
	if v, ok := AnyAsUInt(o.Value); ok {
		return UInt(v)
	}
	return ValueUInt{0, ErrConvertType(o.Value, "uint")}
}

func (o ValueAny) UInt8() ValueUInt8 {
	if o.Err != nil {
		return ValueUInt8{0, o.Err}
	}
	if v, ok := o.Value.(uint8); ok {
		return UInt8(v)
	}
	return ValueUInt8{0, ErrConvertType(o.Value, "uint8")}
}

func (o ValueAny) AsUInt8() ValueUInt8 {
	if o.Err != nil {
		return ValueUInt8{0, o.Err}
	}
	if v, ok := AnyAsUInt8(o.Value); ok {
		return UInt8(v)
	}
	return ValueUInt8{0, ErrConvertType(o.Value, "uint8")}
}

func (o ValueAny) ToUInt8() ValueUInt8 {
	if o.Err != nil {
		return ValueUInt8{0, o.Err}
	}
	if v, ok := AnyAsUInt8(o.Value); ok {
		return UInt8(v)
	}
	return ValueUInt8{0, ErrConvertType(o.Value, "uint8")}
}

func (o ValueAny) UInt16() ValueUInt16 {
	if o.Err != nil {
		return ValueUInt16{0, o.Err}
	}
	if v, ok := o.Value.(uint16); ok {
		return UInt16(v)
	}
	return ValueUInt16{0, ErrConvertType(o.Value, "uint16")}
}

func (o ValueAny) AsUInt16() ValueUInt16 {
	if o.Err != nil {
		return ValueUInt16{0, o.Err}
	}
	if v, ok := AnyAsUInt16(o.Value); ok {
		return UInt16(v)
	}
	return ValueUInt16{0, ErrConvertType(o.Value, "uint16")}
}

func (o ValueAny) ToUInt16() ValueUInt16 {
	if o.Err != nil {
		return ValueUInt16{0, o.Err}
	}
	if v, ok := AnyAsUInt16(o.Value); ok {
		return UInt16(v)
	}
	return ValueUInt16{0, ErrConvertType(o.Value, "uint16")}
}

func (o ValueAny) UInt32() ValueUInt32 {
	if o.Err != nil {
		return ValueUInt32{0, o.Err}
	}
	if v, ok := o.Value.(uint32); ok {
		return UInt32(v)
	}
	return ValueUInt32{0, ErrConvertType(o.Value, "uint32")}
}

func (o ValueAny) AsUInt32() ValueUInt32 {
	if o.Err != nil {
		return ValueUInt32{0, o.Err}
	}
	if v, ok := AnyAsUInt32(o.Value); ok {
		return UInt32(v)
	}
	return ValueUInt32{0, ErrConvertType(o.Value, "uint32")}
}

func (o ValueAny) ToUInt32() ValueUInt32 {
	if o.Err != nil {
		return ValueUInt32{0, o.Err}
	}
	if v, ok := AnyAsUInt32(o.Value); ok {
		return UInt32(v)
	}
	return ValueUInt32{0, ErrConvertType(o.Value, "uint32")}
}

func (o ValueAny) UInt64() ValueUInt64 {
	if o.Err != nil {
		return ValueUInt64{0, o.Err}
	}
	if v, ok := o.Value.(uint64); ok {
		return UInt64(v)
	}
	return ValueUInt64{0, ErrConvertType(o.Value, "uint64")}
}

func (o ValueAny) AsUInt64() ValueUInt64 {
	if o.Err != nil {
		return ValueUInt64{0, o.Err}
	}
	if v, ok := AnyAsUInt64(o.Value); ok {
		return UInt64(v)
	}
	return ValueUInt64{0, ErrConvertType(o.Value, "uint64")}
}

func (o ValueAny) ToUInt64() ValueUInt64 {
	if o.Err != nil {
		return ValueUInt64{0, o.Err}
	}
	if v, ok := AnyAsUInt64(o.Value); ok {
		return UInt64(v)
	}
	return ValueUInt64{0, ErrConvertType(o.Value, "uint64")}
}

func (o ValueAny) Float32() ValueFloat32 {
	if o.Err != nil {
		return ValueFloat32{0, o.Err}
	}
	if v, ok := o.Value.(float32); ok {
		return Float32(v)
	}
	return ValueFloat32{0, ErrConvertType(o.Value, "float32")}
}

func (o ValueAny) AsFloat32() ValueFloat32 {
	if o.Err != nil {
		return ValueFloat32{0, o.Err}
	}
	if v, ok := AnyAsFloat32(o.Value); ok {
		return Float32(v)
	}
	return ValueFloat32{0, ErrConvertType(o.Value, "float32")}
}

func (o ValueAny) ToFloat32() ValueFloat32 {
	if o.Err != nil {
		return ValueFloat32{0, o.Err}
	}
	if v, ok := AnyAsFloat32(o.Value); ok {
		return Float32(v)
	}
	return ValueFloat32{0, ErrConvertType(o.Value, "float32")}
}

func (o ValueAny) Float64() ValueFloat64 {
	if o.Err != nil {
		return ValueFloat64{0, o.Err}
	}
	if v, ok := o.Value.(float64); ok {
		return Float64(v)
	}
	return ValueFloat64{0, ErrConvertType(o.Value, "float64")}
}

func (o ValueAny) AsFloat64() ValueFloat64 {
	if o.Err != nil {
		return ValueFloat64{0, o.Err}
	}
	if v, ok := AnyAsFloat64(o.Value); ok {
		return Float64(v)
	}
	return ValueFloat64{0, ErrConvertType(o.Value, "float64")}
}

func (o ValueAny) ToFloat64() ValueFloat64 {
	if o.Err != nil {
		return ValueFloat64{0, o.Err}
	}
	if v, ok := AnyAsFloat64(o.Value); ok {
		return Float64(v)
	}
	return ValueFloat64{0, ErrConvertType(o.Value, "float64")}
}

func (o ValueAny) MapStringAny() ValueMapStringAny {
	if o.Err != nil {
		return ValueMapStringAny{nil, o.Err}
	}
	if v, ok := o.Value.(map[string]any); ok {
		return ValueMapStringAny{v, nil}
	}
	return ValueMapStringAny{nil, ErrConvertType(o.Value, "map[string]any")}
}

func (o ValueAny) SliceAny() ValueSliceAny {
	if o.Err != nil {
		return ValueSliceAny{nil, o.Err}
	}
	if v, ok := o.Value.([]any); ok {
		return ValueSliceAny{v, nil}
	}
	return ValueSliceAny{nil, ErrConvertType(o.Value, "[]any")}
}
