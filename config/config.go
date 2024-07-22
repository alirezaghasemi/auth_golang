package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		DBName   string `yaml:"dbname"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var cfg Config

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
