package main

import (
	"os"

	"github.com/Katsusan/centaur/internal/commands"
	"github.com/Katsusan/centaur/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	version = "1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "centaur"
	app.Version = version
	app.Usage = "say anything you want about IT career"
	app.EnableBashCompletion = true
	app.Flags = config.GlobalFlags

	app.Commands = []cli.Command{
		commands.StartCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
