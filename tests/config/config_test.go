package config

import (
	"os"
	"testing"

	"../../app/config"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	expectedDefaultCfg := &config.Config{
		AppName:  "GO Rest Api",
		BaseURL:  "http://localhost",
		Port:     8989,
		Database: config.DatabaseConfig{},
	}
	if *cfg != *expectedDefaultCfg { //compare the value
		t.Errorf("Default Config Error found %v\n", cfg)
	}
}

func TestLoadEnvConfig(t *testing.T) {
	os.Setenv("ENV", "production")
	cfg, err := config.LoadEnvConfig("../fixtures/config.json")
	if err != nil {
		t.Errorf("Config Initialization Error %v\n", err)
	}

	expectedAppName := "GO-REST PRODUCTION"
	if cfg.AppName != expectedAppName {
		t.Errorf("Config environment error %v", cfg.AppName)
	}
}

func TestNoEnvPassedShouldReturnDefaultConfig(t *testing.T) {
	cfg, err := config.LoadEnvConfig("../fixtures/config.json")
	if err != nil {
		t.Errorf("Can't load config %v\n", err)
	}
	expectedDefaultCfg := &config.Config{
		AppName:  "GO Rest Api",
		BaseURL:  "http://localhost",
		Port:     8989,
		Database: config.DatabaseConfig{},
	}

	if *cfg != *expectedDefaultCfg {
		t.Errorf("Configuration should same with default configuration! found %v\n", cfg)
	}
}
