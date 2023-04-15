package infra

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	DATABASE = "PG"
	ECHO     = "APP"
	MAMBU    = "MAMBU"
	CORE     = "RPC"
	REDIS    = "REDIS"
	DOKU     = "DOKU"
)

// LoadPgDatabaseCfg load env to new instance of DatabaseCfg.
func LoadPgDatabaseCfg() (*DatabaseCfg, error) {
	var cfg DatabaseCfg

	prefix := DATABASE

	if err := envconfig.Process(prefix, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}

	return &cfg, nil
}

// LoadEchoCfg load env to new instance of EchoCfg.
func LoadEchoCfg() (*AppCfg, error) {
	var cfg AppCfg

	prefix := ECHO

	if err := envconfig.Process(prefix, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}

	return &cfg, nil
}
