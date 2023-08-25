package valconv

import "fmt"

func ErrConvertType(v any, t string) error {
	return fmt.Errorf("convert error, want: %s, got: %T", t, v)
}

func ErrOutOfRange(idx int, c int) error {
	return fmt.Errorf("index %d out of range %d", idx, c)
}
