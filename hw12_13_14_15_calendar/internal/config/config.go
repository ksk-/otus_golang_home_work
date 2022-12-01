package config

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

const (
	defaultLogLevel    = LogLevel(zerolog.DebugLevel)
	defaultStorageType = "memory"
	defaultHost        = "0.0.0.0"
)

type ServiceAddr struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

func (s *ServiceAddr) Addr() string {
	return net.JoinHostPort(s.Host, strconv.Itoa(int(s.Port)))
}

func (s *ServiceAddr) String() string {
	return s.Addr()
}

type DBConfig struct {
	ServiceAddr `yaml:",inline"`
	Database    string `yaml:"database"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
}

func (d *DBConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s", d.User, d.Password, d.Addr(), d.Database)
}

type Logger struct {
	Level  LogLevel `yaml:"level"`
	Pretty bool     `yaml:"pretty"`
}

type Config struct {
	Logger Logger `yaml:"logger"`

	Storage struct {
		Type string   `yaml:"type"`
		DB   DBConfig `yaml:"db"`
	} `yaml:"storage"`

	HTTP ServiceAddr `yaml:"http"`
	GRPC ServiceAddr `yaml:"grpc"`
}

func NewConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
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
