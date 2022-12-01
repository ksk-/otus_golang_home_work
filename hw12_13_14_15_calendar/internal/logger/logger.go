package logger

import (
	"io"
	"os"
	"time"

	"github.com/ksk-/otus_golang_home_work/hw12_13_14_15_calendar/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var global = &Logger{logger: log.Logger}

type Logger struct {
	logger zerolog.Logger
}

func Global() *Logger {
	return global
}

func Debug(msg string) {
	global.Debug(msg)
}

func Info(msg string) {
	global.Info(msg)
}

func Warn(msg string) {
	global.Warn(msg)
}

func Error(msg string) {
	global.Error(msg)
}

func New(cfg *config.Logger) *Logger {
	var logWriter io.Writer
	if cfg.Pretty {
		logWriter = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			FormatTimestamp: func(i interface{}) string {
				parse, _ := time.Parse(time.RFC3339, i.(string))
				return parse.Format("2006-01-02 15:04:05")
			},
		}
	} else {
		logWriter = os.Stdout
	}

	logger := zerolog.New(logWriter).
		Level(zerolog.Level(cfg.Level)).
		With().Timestamp().
		Logger()

	return &Logger{logger: logger}
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

func (l *Logger) WithGlobal() *Logger {
	global = l
	return global
}
