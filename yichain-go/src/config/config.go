package config

import (
	"yichain-go/api"
	"yichain-go/config/simple"
)

type Config struct {
	Log api.Logger
}

func New() (config *Config) {
	config = &Config{Log: simple.New("log.log", "error.log")}
	return config
}
