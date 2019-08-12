package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	BaseURL  string         `json:"base_url"`
	Port     int            `json:"port"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"username"`
	Pass string `json:"password"`
	Db   string `json:"database"`
}

func DefaultConfig() *Config {
	return &Config{
		AppName:  "GO Rest Api",
		BaseURL:  "http://localhost",
		Port:     8989,
		Database: DatabaseConfig{},
	}
}

func LoadEnvConfig() (*Config, error) {
	conf, err := readJsonConfigFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		return nil, err
	}

	env := os.Getenv("ENV")
	if env == "" {
		return DefaultConfig(), nil
	}

	envCfg := conf[env]

	return &envCfg, nil
}

func readJsonConfigFile(path string) (map[string]Config, error) {
	conf := map[string]Config{}
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	byteData, err := ioutil.ReadFile(path)
	jsonErr := json.Unmarshal(byteData, &conf)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return conf, nil
}
