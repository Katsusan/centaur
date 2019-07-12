package config

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/Katsusan/centaur/internal/util"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type Config struct {
	db   *gorm.DB
	parm *Params
}

//Params stores the application config
type Params struct {
	Name           string
	Version        string
	Debug          bool   `yaml:"debug" flag:"debug"`
	LogLevel       string `yaml:"log-level" flag:"log-level"`
	ConfigFile     string
	ConfigPath     string `yaml:"config-path" flag:"config-path"`
	DbServerHost   string `yaml:"db-host" flag:"db-host"`
	DbServerPort   uint16 `yaml:"db-port" flag:"db-port"`
	DbUserName     string `yaml:"db-user" flag:"db-user"`
	DbPassword     string `yaml:"db-password" flag:"db-password"`
	DbName         string `yaml:"db-name" flag:"db-name"`
	HttpServerHost string `yaml:"http-host" flag:"http-host"`
	HttpServerPort uint16 `yaml:"http-port" flag:"http-port"`
	HttpServerMode string `yaml:""http-server-mode flag:"http-server-mode"`
	DaemonMode     bool   `yaml:"daemon" flag:"daemon"`
	DaemonPidPath  string `yaml:"daemon-pid-path" flag:"daemon-pid-path"`
	DaemonLogPath  string `yaml:"daemon-log-path" flag:"daemon-log-path"`
	StaticPath     string `yaml:"static-path" flag:"static-path"`
}

func initLogger(debug bool) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

}

//NewParams will return an instance of struct Params
func NewParams(ctx *cli.Context) *Params {
	params := &Params{}
	params.Name = ctx.App.Name
	params.Version = ctx.App.Version

	return params
}

//NewConfig will return a Config instance with parameters generated from cli.Context
func NewConfig(ctx *cli.Context) *Config {
	initLogger(ctx.GlobalBool(("debug")))

	cfg := &Config{
		parm: NewParams((ctx)),
	}

	return cfg
}

//InitDB will initialize DB configuration
func (c *Config) InitDB(ctx context.Context) error {
	return c.connectToDB(ctx)
}

//connectToDB will establish a new connection to mysql.
func (c *Config) connectToDB(ctx context.Context) error {
	DbDSN := c.DatabaseDSN()
	log.Debugln("will connect to ", DbDSN)

	db, err := gorm.Open("mysql", DbDSN)
	if err != nil {
		log.Fatal(err)
	}

	c.db = db
	return nil
}

//DB will return DB connection.
func (c *Config) DB() *gorm.DB {
	return c.db
}

//CloseDB will cut the connection to DB and return the error if failed.
func (c *Config) CloseDB() error {
	if c.db != nil {
		if err := c.db.Close(); err != nil {
			log.Error("failed to close DB connection")
			return err
		}
		log.Info("DB connection closed")
		c.db = nil
	}
	return nil
}

//ShutDown will release the using resources. inlcuding DB connection, ...
func (c *Config) ShutDown() {
	if err := c.CloseDB(); err != nil {
		return
	}
	log.Info("ShutDown exec OK")
}

//DatabaseDSN will return the connect string of database
func (c *Config) DatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		c.parm.DbUserName,
		c.parm.DbPassword,
		c.parm.DbServerHost,
		c.parm.DbServerPort,
		c.parm.DbName,
	)
}

//DaemonLogFile will return log output file at daemon mode
func (c *Config) DaemonLogFile() string {
	return c.parm.DaemonLogPath
}

//DaemonPidFile will return pid file at daemon mode
func (c *Config) DaemonPidFile() string {
	return c.parm.DaemonPidPath
}

//HttpServerHost will return the host ip which is listening for new connections.
func (c *Config) HttpServerHost() string {
	return c.parm.HttpServerHost
}

//HttpServerPort will return HTTP Server's using port.
func (c *Config) HttpServerPort() uint16 {
	return c.parm.HttpServerPort
}

//HttpServerMode will return under which mode server will be running.
func (c *Config) HttpServerMode() string {
	return c.parm.HttpServerMode
}

//StaticPath will return where are static files stored.
func (c *Config) StaticPath() string {
	return c.parm.StaticPath
}

//ShouldDaemonize will return true if daemon mode is set.
func (c *Config) ShouldDaemonize() bool {
	return c.parm.DaemonMode
}

//Debug will return true if server is set to running at debug mode.
func (c *Config) Debug() bool {
	return c.parm.Debug
}

//LoadFromFile will read config from config file and parse it into *Params
func (p *Params) LoadFromFile(file string) error {
	if !util.FileExist(file) {
		return fmt.Errorf("config file[%s] not found", file)
	}

	yamlconfig, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlconfig, p)

}