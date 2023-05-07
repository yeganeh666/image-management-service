package config

import (
	"image-management-service/pkg/configext"
)

type Config struct {
	*configext.Config
}

func NewConfig(serviceName string) (*Config, error) {
	conf, err := configext.NewConfig("./config/default.yml", serviceName, nil)
	internalConfig := Config{conf}
	return &internalConfig, err
}
