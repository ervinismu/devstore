package main

import (
	"fmt"

	"github.com/ervinismu/devstore/internal/pkg/config"
	"github.com/ervinismu/devstore/internal/pkg/db"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	cfg    config.Config
	DBConn *sqlx.DB
)

func init() {
	// read configuration
	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = configLoad

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
	// init server
	server, err := NewServer(cfg, DBConn)
	if err != nil {
		log.Panic("cannot init server")
	}

	// start server
	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	err = server.Start(appPort)
	if err != nil {
		log.Panic(fmt.Errorf("error cannot start app : %w", err))
	}
}
