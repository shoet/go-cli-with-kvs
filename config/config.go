package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

/*
# redis url
redis://username:password@host:port/database[?[timeout=timeout[d|h|m|s|ms|us|ns]][&clientName=clientName][&libraryName=libraryName] [&libraryVersion=libraryVersion] ]
*/

type Config struct {
	KVSHost           string `env:"KVS_HOST" env_default:"localhost"`
	KVSPort           int    `env:"KVS_PORT" env_default:"6379"`
	KVSPassword       string `env:"KVS_PASSWORD" env_default:""`
	KVSUser           string `env:"KVS_USER" env_default:""`
	KVSExperiationSec int    `env:"KVS_EXPIERATION_SEC" env_default:"300"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return cfg, nil
}
