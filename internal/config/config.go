package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Where to save settings
const configFileName = ".gatorconfig.json"

// What we save in config file
type Config struct {
	DBURL           string `json:"db_url"`            // How to connect to database
	CurrentUserName string `json:"current_user_name"` // Who's logged in
}

// Updates who's logged in
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

// Loads settings from file
func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// Turn JSON into Config struct
	cfg := Config{}
	err = json.NewDecoder(file).Decode(&cfg)
	return cfg, err
}

// Gets full path to config file
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

// Saves settings to file
func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Turn Config struct into JSON
	return json.NewEncoder(file).Encode(cfg)
}
