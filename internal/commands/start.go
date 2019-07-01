package commands

import (
	"github.com/urfave/cli"
)

var startCommand = cli.Command{
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

}
