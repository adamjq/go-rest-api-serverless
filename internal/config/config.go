package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr string `envconfig:"ADDR" default:"8080"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	err = cfg.validate()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) IsLambdaRuntime() bool {
	return os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != ""
}

func (c *Config) validate() error {
	return nil
}
