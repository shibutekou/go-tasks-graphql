package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	HTTP     `yaml:"http"`
	Postgres `yaml:"postgres"`
	Logger   `yaml:"logger"`
}

type HTTP struct {
	Addr            string        `yaml:"addr" env-required:"true"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env-default:"5s"`
}

type Postgres struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"db_name" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
}

type Logger struct {
	Level string `yaml:"level" env-default:"debug"`
}

func Load() *Config {
	configPath := os.Getenv("TODOLIST_CONFIG_PATH")
	if configPath == "" {
		log.Fatal("TODOLIST_CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
