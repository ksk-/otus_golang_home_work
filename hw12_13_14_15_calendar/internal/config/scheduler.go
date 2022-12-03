package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Scheduler struct {
	Logger  Logger        `yaml:"logger"`
	Storage Storage       `yaml:"storage"`
	RMQ     RMQConfig     `yaml:"rmq"`
	Tick    time.Duration `yaml:"tick"`
}

func NewSchedulerConfig(path string) (*Scheduler, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Scheduler
	cfg.Logger.Level = defaultLogLevel
	cfg.Logger.Pretty = false
	cfg.Storage.Type = defaultStorageType
	cfg.Tick = time.Minute

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &cfg, nil
}
