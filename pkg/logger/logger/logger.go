package logger

type Logger interface {
	Info(msg string, params map[string]interface{})
	Infof(format string, args ...interface{})
	Error(msg string, params map[string]interface{})
	Errorf(format string, args ...interface{})
}
