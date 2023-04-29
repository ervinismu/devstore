package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/ervinismu/devstore/internal/pkg/config"
	"github.com/ervinismu/devstore/internal/pkg/db"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	cfg      config.Config
	DBConn   *sqlx.DB
	enforcer *casbin.Enforcer
)

func init() {
	// read configuration
	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = configLoad

	// e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	e, err := casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	if err != nil {
		log.Panic("cannot load app casbin enforcer")
	}
	enforcer = e

	// connect database
	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic("db not established")
	}
	DBConn = db

	// setup logrus
	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)                 // apply log level
	log.SetFormatter(&log.JSONFormatter{}) // define format using json
}

func main() {
	server, err := NewServer(cfg, DBConn, enforcer)
	if err != nil {
		log.Panic("cannot create server")
	}

	serverPort := fmt.Sprintf(":%s", cfg.ServerPort)
	err = server.Start(serverPort)
	if err != nil {
		log.Panic("cannot start server : %s", err)
	}
}
