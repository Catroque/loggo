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
	logc   zerolog.Logger
	level  Level
	active bool
}

func newLogger() Logger {
	//zerolog.SetGlobalLevel(zerolog.InfoLevel)
	return &logger{
		logc: zerolog.New(os.Stdout).With().Timestamp().Logger(),
		active: false,
	}
}

func (log *logger) Debug() Logger {
	log.level = DebugLevel
	if log.logc.GetLevel() <= zerolog.DebugLevel {
		log.active = true
	}
	return log
}

func (log *logger) Msg(message string) {
	if log.active {
		log.logc.Debug().Msg(message)
	}
}
