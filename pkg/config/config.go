package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	// DefaultFilename default config file name
	DefaultFilename string = ".hongbot.json"
)

// Interface config
type Interface interface {
	Save(filepath string) error
	Restore(filepath string) error

	SaveD() error
	RestoreD() error
}

// Config config configification
type Config struct {
	Address  string   `json:"address"`
	Nick     string   `json:"nick"`
	Pass     string   `json:"pass"`
	User     string   `json:"user"`
	Name     string   `json:"name"`
	Channels []string `json:"channels"`
}

// Save config to file
func (c Config) Save(filepath string) error {
	var jsonBytes []byte
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Failed to serialize: %v", err)
	}

	if err := ioutil.WriteFile(filepath, jsonBytes, 0644); err != nil {
		return fmt.Errorf("Failed to writefile: %v", err)
	}

	return nil
}

// SaveD config to default filepath
func (c Config) SaveD() error {
	return c.Save(fmt.Sprintf("%s/%s", os.Getenv("HOME"), DefaultFilename))
}

// Restore config from saved file
func (c *Config) Restore(filepath string) error {
	if _, err := os.Stat(string(filepath)); os.IsNotExist(err) {
		return fmt.Errorf("File not found: %v", err)
	}

	content, err := ioutil.ReadFile(string(filepath))
	if err != nil {
		return fmt.Errorf("Failed to read: %v", err)
	}

	if err := json.Unmarshal(content, c); err != nil {
		return fmt.Errorf("Failed to deserialized: %v", err)
	}

	return nil
}

// RestoreD restore config from default path
func (c *Config) RestoreD() error {
	return c.Restore(fmt.Sprintf("%s/%s", os.Getenv("HOME"), DefaultFilename))
}
