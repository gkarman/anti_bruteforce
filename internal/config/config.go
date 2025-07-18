package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrConfigNotFound  = errors.New("config file not found")
	ErrConfigUnmarshal = errors.New("unable to unmarshal config")
)

type Config struct {
	DBRepo     DBRepo     `yaml:"dbRepo"`
	MemoryRepo MemoryRepo `yaml:"memoryRepo"`
	GrpcServer GrpcServer `yaml:"grpcServer"`
}

type DBRepo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type MemoryRepo struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type GrpcServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func Load(f string) (*Config, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, ErrConfigNotFound
	}

	var c Config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, ErrConfigUnmarshal
	}

	return &c, nil
}
