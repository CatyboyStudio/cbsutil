package logi

type Logi interface {
	Print(msg string)

	Printf(format string, ps ...any)

	Error(err error, msg string)

	Errorf(err error, format string, ps ...any)
}

var Debug Logi = None

var Info Logi = Stdlog

var Warn Logi = Stdlog

var Error Logi = Stdlog
