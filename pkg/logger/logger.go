package logger

import (
	"os"

	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
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
}

type registry struct {
	loge   *zerolog.Event
	level  Level
	active bool
}

func newLogger() Logger {
	//zerolog.SetGlobalLevel(zerolog.InfoLevel)
	return &logger{
		logc: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func newRegistry() *registry {
	return &registry{
		active: false,
	}
}

func (log *logger) Debug() Registry {
	reg := newRegistry()
	reg.level = DebugLevel
	if log.logc.GetLevel() <= zerolog.DebugLevel {
		reg.loge = zl.Debug()
		reg.active = true
	}
	return reg
}

func (reg *registry) Stack() Registry {
	if reg.active {
		reg.loge.Caller()
	}
	return reg
}

func (reg *registry) Msg(message string) {
	if reg.active {
		reg.loge.Msg(message)
	}
}
