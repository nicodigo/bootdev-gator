package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	fullPath, err := getConfigFullPath()
	if err != nil {
		return &Config{}, err
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return &Config{}, err
	}

	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return writeConfig(c)
}

func getConfigFullPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFileName)
	return fullPath, nil
}

func writeConfig(cfg *Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	configPath, err := getConfigFullPath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
