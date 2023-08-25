package logi

type Logi interface {
	Print(msg string)

	Printf(format string, ps ...any)

	Error(err error, msg string)

	Errorf(err error, format string, ps ...any)
}

var LogDebug Logi = None

var LogInfo Logi = Stdlog

var LogWarn Logi = Stdlog

var LogError Logi = Stdlog
