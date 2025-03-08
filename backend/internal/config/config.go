package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"os"
)

type Config struct {
	Env      string `yaml:"env"`
	Database `yaml:"db"`
}

type Database struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH env variable is not set")
	}

	// check if exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Not found %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("Can not read config file: %s", err)
	}
	return &cfg
}
