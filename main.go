package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/exile/commands/server"
	"github.com/ehazlett/exile/version"
)

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "trust orchestrator"
	app.Version = version.FullVersion()
	app.Author = "@ehazlett"
	app.Email = ""
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		return nil
	}
	app.Commands = []cli.Command{
		server.CmdServer,
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
