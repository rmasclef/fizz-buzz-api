package config

import (
	"context"
	"fmt"
	"time"

	"github.com/etf1/go-config"
)

const AppName = "fizzbuzz-api"

var AppVersion = "DEV"

type Config struct {
	// DevMode bool `config:"DEV_MODE"`
	Debug   bool `config:"DEBUG"`

	LogLevel string `config:"LOG_LEVEL"`

	HTTPAddr          string `config:"HTTP_ADDR"`
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func New(context context.Context) *Config {
	cfg := &Config{
		// DevMode: false,
		Debug:   false,

		LogLevel: "INFO",

		HTTPAddr:          ":8001",
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	config.LoadOrFatal(context, cfg)
	fmt.Println(config.TableString(cfg))
	return cfg
}
