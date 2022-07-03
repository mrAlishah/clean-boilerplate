package interfaces

type Logger interface {
	Fatal(msg string, parameters ...interface{})
	Warning(msg string, parameters ...interface{})
	Info(msg string, parameters ...interface{})
}
