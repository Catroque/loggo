package logger

type Logger interface {
	SetLevel(level Level) Logger
    Debug() Registry
    Info()  Registry
    Warn()  Registry
    Error() Registry
    Fatal() Registry
}

func New() Logger {
    return newLogger()
}

type Registry interface {
    Stack() Registry

    Fields(fields map[string]interface{}) Registry

    Msg(message string)
    Msgf(message string, args ...interface{})
}

