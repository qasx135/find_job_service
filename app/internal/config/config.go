package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
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
	PostgreSQL struct {
		Username string `env:"DB_USERNAME" env-required:"true"`
		Password string `env:"DB_PASSWORD" env-required:"true"`
		Host     string `env:"DB_HOST" env-required:"true"`
		Port     string `env:"DB_PORT" env-required:"true"`
		Database string `env:"DB_NAME" env-required:"true"`
	}
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatal(err)
		}
		if err := cleanenv.ReadEnv(instance); err != nil {
			log.Fatal("error initializing config: ", err)
		}
		log.Printf("Config loaded: %+v", instance)

	})
	return instance
}
