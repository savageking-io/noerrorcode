package main

import (
	"fmt"
	"os"

	"github.com/savageking-io/noerrorcode/database"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WebSocket *WebSocketConfig      `yaml:"ws"`
	MySQL     *database.MySQLConfig `yaml:"mysql"`
}

type WebSocketConfig struct {
	Hostname string `yaml:"hostname"`
	Port     uint16 `yaml:"port"`
	URL      string `yaml:"url"`
}

type MongoDBConfig struct {
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	Hostname         string `yaml:"hostname"`
	Port             uint32 `yaml:"port"`
	Database         string `yaml:"database"`
	Retry            bool   `yaml:"retry"`
	RetryAttempts    int    `yaml:"attempts"`
	ReconnectTimeout int    `yaml:"reconnect_timeout"`
}

func (c Config) Read(filepath string) error {
	log.Traceln("Config::Read")
	log.Infof("Reading configuration from %s", filepath)
	buffer, err := os.ReadFile(ConfigFilepath)
	if err != nil {
		return fmt.Errorf("failed to read config: %s", err.Error())
	}

	err = yaml.Unmarshal(buffer, &c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err.Error())
	}

	return nil
}
