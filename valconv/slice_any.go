package valconv

type ValueSliceAny struct {
	Value []any
	Err   error
}

func SliceAny(v []any) ValueSliceAny {
	return ValueSliceAny{v, nil}
}

func (o ValueSliceAny) Len() int {
	return len(o.Value)
}

func (o ValueSliceAny) Get(idx int) ValueAny {
	if o.Err != nil {
		return ValueAny{nil, o.Err}
	}
	if idx >= 0 && idx < len(o.Value) {
		v := o.Value[idx]
		return Any(v)
	} else {
		return ValueAny{nil, ErrOutOfRange(idx, len(o.Value))}
	}
}

func (o ValueSliceAny) Set(idx int, value any) error {
	if o.Err != nil {
		return o.Err
	}
	if idx >= 0 && idx < len(o.Value) {
		o.Value[idx] = value
		return nil
	} else {
		return ErrOutOfRange(idx, len(o.Value))
	}
}

func (o ValueSliceAny) Append(value any) error {
	if o.Err != nil {
		return o.Err
	}
	o.Value = append(o.Value, value)
	return nil
}

func (o ValueSliceAny) RemoveAt(index int) error {
	if o.Err != nil {
		return o.Err
	}
	if index >= 0 && index < len(o.Value) {
		o.Value = append(o.Value[:index], o.Value[index+1:]...)
		return nil
	} else {
		return ErrOutOfRange(index, len(o.Value))
	}
}

func (o ValueSliceAny) RemoveAtClone(index int) error {
	if o.Err != nil {
		return o.Err
	}
	if index >= 0 && index < len(o.Value) {
		tmp := make([]any, len(o.Value)-1)
		copy(tmp, o.Value[:index])
		copy(tmp[index:], o.Value[index+1:])
		o.Value = tmp
		return nil
	} else {
		return ErrOutOfRange(index, len(o.Value))
	}
}
