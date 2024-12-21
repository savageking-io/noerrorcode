package main

import (
	"fmt"
	"os"

	"github.com/savageking-io/noerrorcode/database"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WebSocket *WebSocketConfig        `yaml:"ws"`
	MySQL     *database.MySQLConfig   `yaml:"mysql"`
	MongoDB   *database.MongoDBConfig `yaml:"mongo"`
}

type WebSocketConfig struct {
	Hostname string `yaml:"hostname"`
	Port     uint16 `yaml:"port"`
	URL      string `yaml:"url"`
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

	return nil
}
