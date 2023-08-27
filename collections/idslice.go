package collections

import "slices"

type IdSlice[T any] struct {
	id   int
	Data []T
	ids  []int
}

func (o *IdSlice[T]) Count() int {
	return len(o.Data)
}

func (o *IdSlice[T]) Add(v T) int {
	o.id += 1
	o.Data = append(o.Data, v)
	o.ids = append(o.ids, o.id)
	return o.id
}

func (o *IdSlice[T]) Remove(id int) {
	for i, v := range o.ids {
		if v == id {
			var def T
			o.Data[i] = def
			o.Data = slices.Delete(o.Data, i, i+1)
			o.ids = slices.Delete(o.ids, i, i+1)
			break
		}
	}
}

func (o *IdSlice[T]) Clear() {
	clear(o.Data)
	clear(o.ids)
}
