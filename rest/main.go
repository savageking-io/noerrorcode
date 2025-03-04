package main

import (
	"github.com/savageking-io/noerrorcode/conf"
	//"github.com/savageking-io/noerrorcode/rest/api"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "necrest"
	app.Version = AppVersion
	app.Description = "Smart backend service for smart game developers"
	app.Usage = "REST Microservice of NoErrorCode ecosystem"

	app.Authors = []cli.Author{
		{
			Name:  "savageking.io",
			Email: "i@savageking.io",
		},
		{
			Name:  "Mike Savochkin (crioto)",
			Email: "mike@crioto.com",
		},
	}

	app.Copyright = "2025 (c) savageking.io. All Rights Reserved"

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "Start REST",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config",
					Usage:       "Configuration filepath",
					Value:       ConfigFilepath,
					Destination: &ConfigFilepath,
				},
				cli.StringFlag{
					Name:        "log",
					Usage:       "Specify logging level",
					Value:       LogLevel,
					Destination: &LogLevel,
				},
			},
			Action: Serve,
		},
	}

	_ = app.Run(os.Args)
}
func Serve(c *cli.Context) error {
	if err := SetLogLevel(LogLevel); err != nil {
		log.Errorf("Failed to set log level to %s: %s. Using INFO", LogLevel, err)
	}
	log.Infof("Starting REST service")
	service := new(REST)
	if err := service.Init(conf); err != nil {
		log.Errorf("Failed to init REST service: %s", err.Error())
		return err
	}

	return service.Start()
}

func ReadConfig() error {
	dir, file, err := conf.ExtractDirectoryAndFilenameFromPath(ConfigFilepath)
	if err != nil {
		log.Error("Bad configuration file: %s", err.Error())
		return err
	}

	config := new(conf.Config)
	if err := config.Init(dir); err != nil {
		log.Errorf("Unrecoverable error: %s", err.Error())
		return err
	}
	return nil
}

func SetLogLevel(level string) error {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.SetLevel(log.InfoLevel)
		return err
	}
	log.SetLevel(lvl)
	return nil
}
