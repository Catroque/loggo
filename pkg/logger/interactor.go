package logger

type Logger interface {
	Debug() Registry
}

func New() Logger {
	return newLogger()
}

type Registry interface {
	Stack() Registry
	Msg(string)
}
