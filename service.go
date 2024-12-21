package main

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Serve(c *cli.Context) error {
	SetLogLevel()
	log.Infof("Starting NoErrorCode. Version %s", AppVersion)

	return nil
}

func SetLogLevel() {
	LogLevel = strings.ToLower(LogLevel)
	switch LogLevel {
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
