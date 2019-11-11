package log

import (
	"os"

	"github.com/inconshreveable/log15"
)

type Lazy = log15.Lazy
type Handler = log15.Handler
type Record = log15.Record
type Lvl = log15.Lvl

var (
	LvlCrit  = log15.LvlCrit
	LvlError = log15.LvlError
	LvlWarn  = log15.LvlWarn
	LvlInfo  = log15.LvlInfo
	LvlDebug = log15.LvlDebug
)

type Logger struct {
	log15.Logger
}

func (l *Logger) Trace(msg string, ctx ...interface{}) {
	l.Debug(msg, ctx)
}

func New(ctx ...interface{}) *Logger {
	return &Logger{
		Logger: log15.New(ctx...),
	}
}

func Trace(msg string, ctx ...interface{}) {
	Debug(msg, ctx...)
}

func Debug(msg string, ctx ...interface{}) {
	log15.Root().Debug(msg, ctx...)
}

func Info(msg string, ctx ...interface{}) {
	log15.Root().Info(msg, ctx...)
}

func Warn(msg string, ctx ...interface{}) {
	log15.Root().Warn(msg, ctx...)
}

func Error(msg string, ctx ...interface{}) {
	log15.Root().Error(msg, ctx...)
}

func Crit(msg string, ctx ...interface{}) {
	log15.Root().Crit(msg, ctx...)
}

func SetLogger(verbosity int, json bool) {
	format := TerminalFormatEx()
	if json {
		format = log15.JsonFormat()
	}
	glogger := NewGlogHandler(log15.StreamHandler(os.Stderr, format))
	glogger.Verbosity(log15.Lvl(verbosity))
	log15.Root().SetHandler(glogger)
}
