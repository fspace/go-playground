package db

import (
	"fmt"
	"github.com/spf13/viper"
)

// ... snip ...

type Config struct {
	DatabaseURI string
}

func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("DatabaseURI"),
	}
	if config.DatabaseURI == "" {
		return nil, fmt.Errorf("DatabaseURI must be set")
	}
	return config, nil
}
