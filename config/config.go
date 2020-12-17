package config

import (
	"encoding/json"
	"os"
)

const configPath = "config/config.development.json"

// Config ...
type Config struct {
	AppURL string
	DbURL  string
	DbName string
}

// GetConfig ...
func GetConfig() Config {
	file, err := os.Open(configPath)
	decoder := json.NewDecoder(file)
	var configuration Config
	err = decoder.Decode(&configuration)
	if err != nil {
		return Config{"", "", ""}
	}
	return configuration
}
