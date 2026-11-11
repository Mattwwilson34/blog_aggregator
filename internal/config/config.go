package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configPath = "/internal/config/.config.json"
const DBURL = "postgres://example"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_username"`
}

func (c *Config) Read() (Config, error) {
	filepath, err := getConfigFilepath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) Print() {
	cfg, err := c.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}

func getConfigFilepath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fullConfigPath := cwd + configPath
	return fullConfigPath, nil
}

func Write(cfg *Config) error {
	cfgJson, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	filepath, err := getConfigFilepath()
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath, cfgJson, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username
	err := Write(cfg)
	if err != nil {
		return err
	}
	return nil
}
