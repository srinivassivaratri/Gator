// This package handles configuration file operations
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Name of our config file that will be stored in user's home directory
const configFileName = ".gatorconfig.json"

// Config struct defines the structure of our JSON config file
// The `json:` tags tell the JSON encoder/decoder what names to use in the JSON file
type Config struct {
	DBURL           string `json:"db_url"`            // Database connection URL
	CurrentUserName string `json:"current_user_name"` // Current user's name
}

// SetUser is a method on Config that updates the username and saves to file
func (cfg *Config) SetUser(username string) error {
	// Update the username in our config struct
	cfg.CurrentUserName = username
	// Write the updated config to file
	return write(*cfg)
}

func Read() (Config, error) {
	// Get the full path to the config file (e.g., /home/user/.gatorconfig.json)
	fullPath, err := getConfigFilePath()
	if err != nil {
		// If we can't get the path, return an empty config and the error
		return Config{}, err
	}

	// Open the config file for reading
	// os.Open returns a file handle that we can read from
	file, err := os.Open(fullPath)
	if err != nil {
		// If we can't open the file (maybe it doesn't exist), return empty config and error
		return Config{}, err
	}
	// Make sure we close the file when we're done with it
	// defer means this will run at the end of the function
	defer file.Close()

	// Create a new JSON decoder that will read from our file
	// This is more efficient than reading the whole file into memory first
	decoder := json.NewDecoder(file)

	// Create an empty Config struct that we'll fill with the JSON data
	cfg := Config{}

	// Try to decode the JSON into our Config struct
	// This will match the JSON fields to our struct fields using the json tags
	err = decoder.Decode(&cfg)
	if err != nil {
		// If we can't decode the JSON (maybe it's malformed), return empty config and error
		return Config{}, err
	}

	// Everything worked! Return the filled Config struct and nil error
	return cfg, nil
}

// getConfigFilePath builds the full path to our config file
func getConfigFilePath() (string, error) {
	// Get the user's home directory (e.g., /home/username on Linux)
	home, err := os.UserHomeDir()
	if err != nil {
		// If we can't get home directory, return error
		return "", err
	}
	// Join home directory path with our config filename
	// e.g., /home/username/.gatorconfig.json
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

// write saves the config struct to the JSON file
func write(cfg Config) error {
	// Get the full path where we'll save the config
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Create (or truncate) the config file
	// os.Create will create the file if it doesn't exist or clear it if it does
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	// Make sure we close the file when we're done
	defer file.Close()

	// Create a JSON encoder that will write to our file
	encoder := json.NewEncoder(file)
	// Encode our config struct to JSON and write it to the file
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
