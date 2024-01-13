package config

import (
	"encoding/json"
	"os"

	"github.com/zsandibe/basic-marketplace/pkg"
)

type Config struct {
	Port    string
	Migrate string
	DB      struct {
		DSN    string
		Driver string
	}
}

func NewConfig() (Config, error) {
	configFile, err := os.Open("config.yml")
	if err != nil {
		pkg.ErrorLog.Printf("Can`t open config.yml: %v", err)
		return Config{}, err
	}
	defer configFile.Close()
	var config Config

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		pkg.ErrorLog.Printf("Can`t decode config.yml: %v", err)
		return Config{}, err
	}
	return config, nil
}
