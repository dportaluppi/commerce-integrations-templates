package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoURI string `default:"mongodb://localhost:27017" split_words:"true"`
}

func Load() (*Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
