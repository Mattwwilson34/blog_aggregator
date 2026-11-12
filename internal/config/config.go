package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = ".gatorconfig.json"
const DBURL = "postgres://example"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_username"`
}

func Read() (Config, error) {
	filepath, err := getConfigFilepath()
	fmt.Fprintln(os.Stderr, filepath)
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

func Print() {
	cfg, err := Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	_, err = fmt.Fprintln(os.Stdout, string(b))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

	}
}

func getConfigFilepath() (string, error) {
	// cwd, err := os.Getwd()
	cwd, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullConfigPath := cwd + "/" + fileName
	return fullConfigPath, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username
	err := Write(cfg)
	if err != nil {
		return err
	}
	return nil
}
