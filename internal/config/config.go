package config

import "github.com/caarlos0/env/v8"

type Config struct {
	TelegramAPIToken string `env:"TELEGRAM_API_TOKEN,notEmpty,unset,file"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
