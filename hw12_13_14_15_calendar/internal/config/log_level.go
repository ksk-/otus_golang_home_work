package config

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

type LogLevel zerolog.Level

const (
	DebugLogLevel = LogLevel(zerolog.DebugLevel)
	InfoLogLevel  = LogLevel(zerolog.InfoLevel)
	WarnLogLevel  = LogLevel(zerolog.WarnLevel)
	ErrorLogLevel = LogLevel(zerolog.ErrorLevel)
)

func (l *LogLevel) UnmarshalYAML(value *yaml.Node) error {
	level, ok := levels[strings.ToLower(value.Value)]
	if !ok {
		return fmt.Errorf(`unknown log level "%s"`, value.Value)
	}
	*l = level
	return nil
}

var levels = map[string]LogLevel{
	"debug": DebugLogLevel,
	"info":  InfoLogLevel,
	"warn":  WarnLogLevel,
	"error": ErrorLogLevel,
}
