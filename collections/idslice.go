package collections

import "slices"

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

func (o IdSlice) Remove(id int) {
	for i, v := range o.ids {
		if v == id {
			slices.Delete(o.Data, i, i+1)
			slices.Delete(o.ids, i, i+1)
			break
		}
	}
}
