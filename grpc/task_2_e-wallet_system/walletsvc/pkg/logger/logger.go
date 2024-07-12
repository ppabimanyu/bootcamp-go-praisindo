package logger

type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Error(string, ...interface{})
	Warn(string, ...interface{})
	Type() string
	Default() interface{}
}
