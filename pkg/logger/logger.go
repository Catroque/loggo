package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	NoLevel
	Disabled
	TraceLevel
)

type logger struct {
	logc  zerolog.Logger
	level Level
}

func newLogger() *logger {

	return &logger{
		logc: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (log *logger) Debug() *logger {
	log.level = logger.DebugLevel
	return log
}


func (log *logger) Msg(message string) {
	logc.Debug().Msg(message)
}

