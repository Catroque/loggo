package logger

type Logger interface {
	Debug() Logger
}

func New() Logger {
	return newLogger()
}

