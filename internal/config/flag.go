package config

import (
	"errors"
	"flag"
)

var (
	ErrInvalidConfigPath = errors.New("please provide valid path to config")
)

type FlagConfig struct {
	PathToConfigFile string
}

func NewFlagConfig() *FlagConfig {
	return &FlagConfig{}
}

func (cfg *FlagConfig) Parse() error {
	flag.StringVar(&cfg.PathToConfigFile, "config", "", "/path/to/config")
	flag.Parse()

	if cfg.PathToConfigFile == "" {
		return ErrInvalidConfigPath
	}

	return nil
}
