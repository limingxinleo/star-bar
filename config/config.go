package config

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed .env.json
var configJson []byte

type Config struct {
	Repo  string `json:"repo"`
	Token string `json:"token"`
}

func Init() *Config {
	config := &Config{}
	err := json.Unmarshal(configJson, config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return config
}
