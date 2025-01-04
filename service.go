package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/savageking-io/noerrorcode/database"
	"github.com/savageking-io/noerrorcode/steam"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type NoErrorCode struct {
	//Steam     *steam.Steam
	//Config    *Config
	//WebSocket *WebSocket
	// MySQL     *database.MySQL
	// Mongo     *database.MongoDB
}

var nec *NoErrorCode

func Serve(c *cli.Context) error {
	SetLogLevel(LogLevel)
	log.Infof("Starting NoErrorCode. Version %s", AppVersion)

	//nec = new(NoErrorCode)

	config := new(Config)
	if err := config.Read(ConfigFilepath); err != nil {
		log.Errorf("Failed to read config: %s", err.Error())
		return err
	}

	mysql, err := InitMySQL(config.MySQL)
	if err != nil {
		log.Errorf("MySQL Init failed: %s", err.Error())
		return err
	}
	mysql.AutoMigrate()
	mysql.PopulateIfFresh()

	mongo, err := InitMongo(config.MongoDB)
	if err != nil {
		log.Errorf("MongoDB Init failed: %s", err.Error())
		return err
	}

	steam := new(steam.Steam)
	if err := steam.Init(config.Steam, ""); err != nil {
		log.Errorf("Steam Init failed: %s", err.Error())
		return err
	}

	clients := new(ClientManager)
	if err := clients.Init(steam, mysql, mongo, config.Crypto); err != nil {
		log.Errorf("Failed to initialize Client Manager: %s", err.Error())
		return err
	}

	ws := new(WebSocket)
	if err := ws.Init(config.WebSocket, clients); err != nil {
		log.Errorf("Failed to initialize WebSocket: %s", err.Error())
		return err
	}

	characterManager := new(CharacterManager)
	if err := characterManager.Init(mysql); err != nil {
		log.Errorf("Character Manager Init failed: %s", err.Error())
		return err
	}

	return ws.Run()
}

func InitMongo(config *database.MongoDBConfig) (*database.MongoDB, error) {
	Mongo := new(database.MongoDB)
	if config == nil {
		return nil, fmt.Errorf("nil mongo config")
	}
	if err := Mongo.Init(config); err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	connectTime := time.Unix(0, 0)
	attempt := 0
	for {
		if time.Since(connectTime) > time.Duration(time.Duration(config.ReconnectTimeout)*time.Second) {
			connectTime = time.Now()
			attempt++
			if err := Mongo.Connect(); err != nil {
				log.Errorf("%s", err.Error())
				if !config.Retry {
					return nil, fmt.Errorf("mongo init: %s", err.Error())
				}
				if config.RetryAttempts > 0 && attempt >= config.RetryAttempts {
					return nil, fmt.Errorf("mongo init: failed after %d attempts: %s", attempt, err.Error())
				}
				log.Infof("Will try to reconnect in %d seconds", config.ReconnectTimeout)
			} else {
				break
			}
		}

		time.Sleep(time.Millisecond * 100)
	}

	return Mongo, nil
}

func InitMySQL(config *database.MySQLConfig) (*database.MySQL, error) {
	MySQL := new(database.MySQL)

	if config == nil {
		return nil, fmt.Errorf("mysql init: nil config")
	}

	if err := MySQL.Init(config); err != nil {
		return nil, fmt.Errorf("mysql init: %s", err.Error())
	}

	connectTime := time.Unix(0, 0)
	attempt := 0
	for {
		if time.Since(connectTime) > time.Duration(time.Duration(config.RetryTimeout)*time.Second) {
			connectTime = time.Now()
			attempt++
			if err := MySQL.Connect(); err != nil {
				log.Errorf("%s", err.Error())
				if !config.Retry {
					return nil, fmt.Errorf("mysql init: %s", err.Error())
				}
				if config.RetryAttempts > 0 && attempt >= config.RetryAttempts {
					return nil, fmt.Errorf("mysql init: failed after %d sttempts: %s", attempt, err.Error())
				}
				log.Infof("Will try to reconnect in %d seconds", config.RetryTimeout)
			} else {
				break
			}
		}

		time.Sleep(time.Millisecond * 100)
	}

	return MySQL, nil
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
