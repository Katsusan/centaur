package commands

import (
	"context"
	"os"
	"strconv"
	"syscall"

	"github.com/Katsusan/centaur/internal/config"
	"github.com/Katsusan/centaur/internal/util"
	daemon "github.com/sevlyar/go-daemon"
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

	if err := conf.InitDB(cctx); err != nil {
		log.Fatal(err)
	}

	dcxt := &daemon.Context{
		Args:        ctx.Args(),
		LogFileName: conf.DaemonLogFile(),
		LogFilePerm: 0640,
		PidFileName: conf.DaemonPidFile(),
		PidFilePerm: 0644,
	}

	if !daemon.WasReborn() && conf.ShouldDaemonize() {
		conf.ShutDown()
		cancel()

		if pid, ok := childAlreadyRunning(conf.DaemonPidFile()); ok {
			log.Infof("Daemon already running with PID[%d]\n", pid)
			return
		}

		child, err := dcxt.Reborn()
		if err != nil {
			log.Fatalf("failed to create a new process, %v", err)
		}

		if child != nil {
			if err := util.OverWrite(conf.DaemonPidFile(), []byte(strconv.Itoa(child.Pid))); err != nil {
				log.Fatalf("failed to write to pid file, [%v]\n", err)
			}

			log.Infof("Daemon started with Pid(%d)\n", child.Pid)
			return
		}
	}

	log.Infof("HTTP Server started at %s:%d\n", conf.HttpServerHost(), conf.HttpServerPort())

}

//childAlreadyRunning will try to find  process corresponding to path file
//and return its pid and running state.
func childAlreadyRunning(path string) (pid int, running bool) {
	if !util.FileExist(path) {
		return pid, false
	}

	pid, err := daemon.ReadPidFile(path)
	if err != nil {
		return pid, false
	}

	proc, err := os.FindProcess(int(pid))
	if err != nil {
		return pid, false
	}

	return pid, proc.Signal(syscall.Signal(0)) == nil
}
