package commands

import (
	"context"

	"github.com/Katsusan/centaur/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "start backend web server",
	Flags:  startFlags,
	Action: startAction,
}

var startFlags = []cli.Flag{
	cli.IntFlag{
		Name:   "port, p",
		Usage:  "backend server port",
		Value:  80,
		EnvVar: "CENTAUR_HTTP_PORT",
	},
	cli.StringFlag{
		Name:   "host, h",
		Usage:  "backend server host ip",
		Value:  "",
		EnvVar: "CENTAUR_HTTP_HOST",
	},
	cli.BoolFlag{
		Name:   "daemonize, d",
		Usage:  "run CENTAUR at Daemon mode",
		EnvVar: "CENTAUR_DAEMON_MODE",
	},
}

func startAction(ctx *cli.Context) {
	cctx, cancel := context.WithCancel(context.Background())

	conf := config.NewConfig(ctx)
	if conf.HttpServerPort() < 0 || conf.HttpServerPort() > 65535 {
		log.Fatal("Server port must be an integer between 0~65535")
	}

	if err := config.InitDB(cctx); err != nil {
		log.Fatal(err)
	}
}
