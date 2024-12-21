package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "noerrorcode"
	app.Version = AppVersion
	app.Description = "Smart backend service for smart game developers"
	app.Usage = "Core service for NoErrorCode solution"

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

	app.Copyright = "2025 (c) savageking.io. All Right Reserved"

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "Start core service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config",
					Usage:       "Path to config file",
					Value:       "noerrorcode.yaml",
					Destination: &ConfigFilepath,
				},
				cli.StringFlag{
					Name:        "log",
					Usage:       "Specify logging level",
					Value:       "info",
					Destination: &LogLevel,
				},
			},
			Action: Serve,
		},
	}

	app.Run(os.Args)
}
