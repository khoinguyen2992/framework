package logs

type Logger interface {
	Debug(args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
}
