package main

import (
	"strings"

	"github.com/savageking-io/noerrorcode/database"
	"github.com/savageking-io/noerrorcode/steam"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type NoErrorCode struct {
	Steam     *steam.Steam
	Config    *Config
	WebSocket *WebSocket
	MySQL     *database.MySQL
	Mongo     *database.MongoDB
}

var nec *NoErrorCode

func Serve(c *cli.Context) error {
	SetLogLevel(LogLevel)
	log.Infof("Starting NoErrorCode. Version %s", AppVersion)

	nec = new(NoErrorCode)

	nec.Config = new(Config)
	if err := nec.Config.Read(ConfigFilepath); err != nil {
		log.Errorf("Failed to read config: %s", err.Error())
		return err
	}

	nec.WebSocket = new(WebSocket)
	if err := nec.WebSocket.Init(nec.Config.WebSocket); err != nil {
		log.Errorf("Failed to initialize WebSocket: %s", err.Error())
		return err
	}

	nec.Steam = new(steam.Steam)
	if err := nec.Steam.Init(nec.Config.Steam); err != nil {
		log.Errorf("Steam Init failed: %s", err.Error())
		return err
	}

	nec.MySQL = new(database.MySQL)
	if err := nec.MySQL.Init(nec.Config.MySQL); err != nil {
		log.Errorf("MySQL Init failed: %s", err.Error())
		return err
	}
	if err := nec.MySQL.Connect(); err != nil {
		log.Errorf("MySQL Connect failed: %s", err.Error())
		return err
	}
	nec.MySQL.AutoMigrate()
	nec.MySQL.PopulateIfFresh()

	nec.Mongo = new(database.MongoDB)
	if err := nec.Mongo.Init(nec.Config.MongoDB); err != nil {
		log.Errorf("MongoDB Init failed: %s", err.Error())
		return err
	}

	characterManager := new(CharacterManager)
	if err := characterManager.Init(nec.MySQL); err != nil {
		log.Errorf("Character Manager Init failed: %s", err.Error())
		return err
	}

	return nec.WebSocket.Run()
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
