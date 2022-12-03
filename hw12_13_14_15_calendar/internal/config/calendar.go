package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Calendar struct {
	Logger  Logger      `yaml:"logger"`
	Storage Storage     `yaml:"storage"`
	HTTP    ServiceAddr `yaml:"http"`
	GRPC    ServiceAddr `yaml:"grpc"`
}

func NewCalendarConfig(path string) (*Calendar, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Calendar
	cfg.Logger.Level = defaultLogLevel
	cfg.Logger.Pretty = false
	cfg.Storage.Type = defaultStorageType
	cfg.HTTP.Host = defaultHost
	cfg.GRPC.Host = defaultHost

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
