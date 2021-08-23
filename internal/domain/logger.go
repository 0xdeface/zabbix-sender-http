package domain

type LogLevel int

const (
	DEBUG = iota
	INFO
	ERROR
)

type Logger interface {
	Log(level LogLevel, message string)
}

type LoggerConstructor interface {
	New(level Logger) Logger
}
