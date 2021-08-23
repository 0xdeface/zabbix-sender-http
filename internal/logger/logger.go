package logger

import (
	"fmt"
	"zabbix-http/internal/domain"
)

type Logger struct {
	level domain.LogLevel
}

func (l Logger) Log(level domain.LogLevel, message string) {
	if level >= l.level {
		fmt.Println(message)
	}
}
func NewLogger(level domain.LogLevel) *Logger {
	return &Logger{level: level}
}
