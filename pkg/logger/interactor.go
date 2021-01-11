package logger

type Logger interface {
	Debug()
}

func New() Logger {
	return newLogger()
}

