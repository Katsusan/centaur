package config

import (
	"github.com/urfave/cli"
)

//GlobalFlags are Global CLI Flags
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "run in debug mode",
		EnvVar: "CENTAUR_DEBUG",
	},
	cli.StringFlag{
		Name:   "log-level",
		Usage:  "specify log level, value is one of [trace, debug, info, warning, error, fatal or panic]",
		Value:  "info",
		EnvVar: "CENTAUR_LOG_LEVEL",
	},
	cli.StringFlag{
		Name:   "config-file, c",
		Usage:  "load configuration from `FILENAME`",
		Value:  "./centaur.yml",
		EnvVar: "CENTAUR_CONFIG_FILE",
	},
	cli.StringFlag{
		Name:   "http-host, h",
		Usage:  "HTTP server host",
		Value:  "",
		EnvVar: "CENTAUR_HTTP_HOST",
	},
	cli.StringFlag{
		Name:   "http-port, p",
		Usage:  "HTTp server port",
		Value:  "80",
		EnvVar: "CENTAUR_HTTP_PORT",
	},
	cli.StringFlag{
		Name:   "mysql-host",
		Usage:  "MySQL server host",
		Value:  "",
		EnvVar: "CENTAUR_MYSQL_HOST",
	},
	cli.StringFlag{
		Name:   "mysql-port",
		Usage:  "MySQL server host",
		Value:  "3306",
		EnvVar: "CENTAUR_MYSQL_HOST",
	},
	cli.BoolFlag{
		Name:   "daemonize, d",
		Usage:  "run centaur as Daemon",
		EnvVar: "CENTAUR_DAEMON_MODE",
	},
}
