package main

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Serve(c *cli.Context) error {
	SetLogLevel(LogLevel)
	log.Infof("Starting NoErrorCode. Version %s", AppVersion)

	config := new(Config)
	if err := config.Read(ConfigFilepath); err != nil {
		log.Errorf("Failed to read config: %s", err.Error())
		return err
	}

	ws := new(WebSocket)
	if err := ws.Init(config.WebSocket); err != nil {
		log.Errorf("Failed to initialize WebSocket: %s", err.Error())
		return err
	}

	return ws.Run()
}

func SetLogLevel(level string) {
	level = strings.ToLower(level)
	switch level {
	case "trace":
		log.SetLevel(log.TraceLevel)
		return
	case "debug":
		log.SetLevel(log.DebugLevel)
		return
	case "warn":
		log.SetLevel(log.WarnLevel)
		return
	case "error":
		log.SetLevel(log.ErrorLevel)
		return
	case "fatal":
		log.SetLevel(log.FatalLevel)
		return
	default:
		log.SetLevel(log.InfoLevel)
	}
}
