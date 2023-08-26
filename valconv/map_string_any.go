package valconv

type ValueMapStringAny = ValueMapStringT[map[string]any, any]

type ValueMapStringT[S ~map[string]T, T any] struct {
	Value S
	Err   error
}

func MapStringAny(v map[string]any) ValueMapStringAny {
	return MapStringT(v)
}

func MapStringT[S ~map[string]T, T any](v S) ValueMapStringT[S, T] {
	return ValueMapStringT[S, T]{v, nil}
}

func (o ValueMapStringT[S, T]) Get(n string) ValueAny {
	if o.Err != nil {
		return ValueAny{nil, o.Err}
	}
	if o.Value != nil {
		if v, ok := o.Value[n]; ok {
			return Any(v)
		}
	}
	return Any(nil)
}

func (o ValueMapStringT[S, T]) GetData(n string, def T) T {
	if o.Err != nil {
		return def
	}
	if o.Value != nil {
		if v, ok := o.Value[n]; ok {
			return v
		}
	}
	return def
}

func (o ValueMapStringT[S, T]) Set(n string, value T) error {
	if o.Err == nil {
		return o.Err
	}
	if o.Value == nil {
		o.Value = make(S)
	}
	o.Value[n] = value
	return nil
}

func (o ValueMapStringT[S, T]) Delete(n string) error {
	if o.Err != nil {
		return o.Err
	}
	if o.Value != nil {
		delete(o.Value, n)
	}
	return nil
}
