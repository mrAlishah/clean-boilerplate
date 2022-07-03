package interfaces

type Logger interface {
	Fatal(msg string, parameters ...interface{})
	Warring(msg string, parameters ...interface{})
	Info(msg string, parameters ...interface{})
}
