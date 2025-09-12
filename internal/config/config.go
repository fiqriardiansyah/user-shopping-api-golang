package config

import (
	"os"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/helper"
)

func NewConfig() *helper.Config {
	cfg := &helper.Config{
		Prefix:  os.Getenv("PREFIX"),
		BaseUrl: os.Getenv("BASE_URL"),
	}

	if cfg.BaseUrl == "" {
		panic("Base url not found")
	}

	return cfg
}
