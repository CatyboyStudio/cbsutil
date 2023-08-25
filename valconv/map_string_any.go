package valconv

type ValueMapStringAny struct {
	Value map[string]any
	Err   error
}

func MapStringAny(v map[string]any) ValueMapStringAny {
	return ValueMapStringAny{v, nil}
}

func (o ValueMapStringAny) Get(n string) ValueAny {
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

func (o ValueMapStringAny) Set(n string, value any) error {
	if o.Err == nil {
		return o.Err
	}
	if o.Value == nil {
		o.Value = make(map[string]any)
	}
	o.Value[n] = value
	return nil
}

func (o ValueMapStringAny) Delete(n string) error {
	if o.Err != nil {
		return o.Err
	}
	if o.Value != nil {
		delete(o.Value, n)
	}
	return nil
}
