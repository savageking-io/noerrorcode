package main

import (
	"fmt"
	"os"

	"github.com/savageking-io/noerrorcode/database"
	"github.com/savageking-io/noerrorcode/steam"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WebSocket *WebSocketConfig        `yaml:"ws"`
	MySQL     *database.MySQLConfig   `yaml:"mysql"`
	MongoDB   *database.MongoDBConfig `yaml:"mongo"`
	Steam     *steam.Config           `yaml:"steam"`
	Crypto    *CryptoConfig           `yaml:"crypto"`
}

type WebSocketConfig struct {
	Hostname string `yaml:"hostname"`
	Port     uint16 `yaml:"port"`
	URL      string `yaml:"url"`
}

// CryptoConfig - configuration for JWT Tokens
// Currently only HS256 is supported
type CryptoConfig struct {
	Key string `yaml:"sign_key"`
}

func (c *Config) Read(filepath string) error {
	log.Traceln("Config::Read")
	log.Infof("Reading configuration from %s", filepath)
	buffer, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read config: %s", err.Error())
	}

	err = yaml.Unmarshal(buffer, c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err.Error())
	}

	if c.Crypto == nil || c.Crypto.Key == "" {
		return fmt.Errorf("crypto key not set")
	}

	return nil
}
