package config

import (
	"os"
)

type Config interface {
	Token() string
}

type UserConfig struct {}

func (c UserConfig) Token() string {
	// fetch token from config file in $HOME
	return os.Getenv("GITHUB_PERSONAL_TOKEN")
}

// add handler that accepts a string and stores into config file in $HOME
