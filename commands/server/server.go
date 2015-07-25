package server

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ehazlett/exile/api"
	"github.com/ehazlett/exile/config"
)

var CmdServer = cli.Command{
	Name:  "server",
	Usage: "Start API server",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "listen, l",
			Value: ":8080",
			Usage: "Listen address for API",
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: "",
			Usage: "Path to config",
		},
	},
	Action: startServer,
}

func startServer(c *cli.Context) {
	configPath := c.String("config")
	if configPath == "" {
		log.Fatal("you must specify a config file")
	}
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("unable to load config: %s", err)
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	signers, err := config.LoadSigners(cfg)
	if err != nil {
		log.Fatal(err)
	}

	addr := c.String("listen")

	a, err := api.NewAPI(addr, signers)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
