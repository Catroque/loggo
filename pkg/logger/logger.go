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
		loge: nil,
		level: NoLevel,
		active: false,
	}
}

func (reg *registry) setLevel(log *logger, level Level) *registry {
	if log.logc.GetLevel() <= zerolog.DebugLevel {
		reg.loge   = zl.Debug()
		reg.level  = level
		reg.active = true
	}
	return reg
}


func (log *logger) Debug() Registry {
	reg := newRegistry().setLevel(log, DebugLevel)
	return reg
}

func (log *logger) Info() Registry {
	reg := newRegistry().setLevel(log, InfoLevel)
	return reg
}

func (log *logger) Warn() Registry {
	reg := newRegistry().setLevel(log, WarnLevel)
	return reg
}

func (log *logger) Error() Registry {
	reg := newRegistry().setLevel(log, ErrorLevel)
	return reg
}

func (log *logger) Fatal() Registry {
	reg := newRegistry().setLevel(log, FatalLevel)
	return reg
}


func (reg *registry) Stack() Registry {
	if reg.active {
		reg.loge.Caller()
	}
	return reg
}

func (reg *registry) Fields(fields map[string]interface{}) Registry {
	if reg.active {
		reg.loge.Fields(fields)
	}
	return reg
}

func (reg *registry) Msg(message string) {
	if reg.active {
		reg.loge.Msg(message)
	}
}

func (reg *registry) Msgf(message string, args ...interface{}) {
	if reg.active {
		reg.loge.Msgf(message, args...)
	}
}
