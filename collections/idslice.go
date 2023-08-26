package collections

type IdSlice struct {
	id   int
	Data []any
	ids  []int
}

func (o *IdSlice) Add(v any) int {
	o.id += 1
	o.Data = append(o.Data, v)
	o.ids = append(o.ids, o.id)
	return o.id
}

func (o *IdSlice) Remove(id int) {
	for i, v := range o.ids {
		if v == id {
			copy(o.Data[i:], o.Data[i+1:])
			o.Data[len(o.Data)] = nil
			o.Data = o.Data[:len(o.Data)-1]
			copy(o.ids[i:], o.ids[i+1:])
			o.ids = o.ids[:len(o.ids)-1]
			break
		}
	}
}
