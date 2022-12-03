package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Sender struct {
	Logger Logger    `yaml:"logger"`
	RMQ    RMQConfig `yaml:"rmq"`
}

func NewSenderConfig(path string) (*Sender, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Sender
	cfg.Logger.Level = defaultLogLevel
	cfg.Logger.Pretty = false

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
