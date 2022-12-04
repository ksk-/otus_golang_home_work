package config

import (
	"fmt"
	"net"
	"strconv"

	"github.com/rs/zerolog"
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

type Storage struct {
	Type string    `yaml:"type"`
	DB   *DBConfig `yaml:"db"`
}

type RMQConfig struct {
	ServiceAddr `yaml:",inline"`
	Scheme      string `yaml:"scheme"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Queue       string `yaml:"queue"`
}

func (r *RMQConfig) URI() string {
	return fmt.Sprintf("%s://%s:%s@%s", r.Scheme, r.User, r.Password, r.Addr())
}
