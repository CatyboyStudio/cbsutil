package logi

type none int

// Error implements Logi.
func (none) Error(err error, msg string) {
}

// Errorf implements Logi.
func (none) Errorf(err error, format string, ps ...any) {
}

// Print implements Logi.
func (none) Print(msg string) {
}

// Printf implements Logi.
func (none) Printf(format string, ps ...any) {
}

var None Logi = none(0)
