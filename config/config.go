package config

import (
	"encoding/json"
	"os"
)

// Config ...
type Config struct {
	AppURL string
	DbURL  string
	DbName string
}

// GetConfig ...
func GetConfig(env string) Config {
	file, err := os.Open("config/config." + env + ".json")
	decoder := json.NewDecoder(file)
	var configuration Config
	err = decoder.Decode(&configuration)
	if err != nil {
		return Config{"", "", ""}
	}
	return configuration
}
