package logger

type Logger interface {
	Debug() Logger
	Msg(string)
}

func New() Logger {
	return newLogger()
}

