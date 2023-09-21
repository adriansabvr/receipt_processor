package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/palantir/stacktrace"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"level"   env:"LOG_LEVEL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, stacktrace.Propagate(err, "failed to read config file")
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, stacktrace.Propagate(err, "failed to read env variables")
	}

	return cfg, nil
}
