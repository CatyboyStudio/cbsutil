package valconv

import "slices"

type ValueSliceAny = ValueSliceT[[]any, any]

func SliceAny(v []any) ValueSliceAny {
	return SliceT(v)
}

type ValueSliceT[S ~[]T, T any] struct {
	Value S
	Err   error
}

func SliceT[S ~[]T, T any](v S) ValueSliceT[S, T] {
	return ValueSliceT[S, T]{v, nil}
}

func (o ValueSliceT[S, T]) Len() int {
	return len(o.Value)
}

func (o ValueSliceT[S, T]) Get(idx int) ValueAny {
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

func (o ValueSliceT[S, T]) Set(idx int, value T) error {
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

func (o *ValueSliceT[S, T]) Append(value T) error {
	if o.Err != nil {
		return o.Err
	}
	o.Value = append(o.Value, value)
	return nil
}

func (o *ValueSliceT[S, T]) RemoveAt(index int) error {
	if o.Err != nil {
		return o.Err
	}
	if index >= 0 && index < len(o.Value) {
		slices.Delete(o.Value, index, index+1)
		return nil
	} else {
		return ErrOutOfRange(index, len(o.Value))
	}
}

func (o *ValueSliceT[S, T]) RemoveAtClone(index int) error {
	if o.Err != nil {
		return o.Err
	}
	if index >= 0 && index < len(o.Value) {
		tmp := slices.Clone(o.Value)
		slices.Delete(tmp, index, index+1)
		o.Value = tmp
		return nil
	} else {
		return ErrOutOfRange(index, len(o.Value))
	}
}
