package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug  bool   `env:"IS_DEBUG" env-default:"true"`
	IsProd   bool   `env:"IS_PROD" env-default:"false"`
	LogLevel string `env:"LOG_LEVEL" env-default:"debug"`
	Listen   struct {
		Host string `env:"LISTEN_HOST" env-default:"0.0.0.0"`
		Port string `env:"LISTEN_PORT" env-default:"10100"`
	}
	AdminUser struct {
		Username string `env:"ADMIN_USERNAME" env-default:"admin"`
		Password string `env:"ADMIN_PASSWORD" env-default:"admin"`
	}
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			log.Fatal("error initializing config: ", err)
		}

	})
	return instance
}
