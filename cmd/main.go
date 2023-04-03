package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/ervinismu/devstore/internal/app/controller"
	"github.com/ervinismu/devstore/internal/app/repository"
	"github.com/ervinismu/devstore/internal/app/service"
	"github.com/ervinismu/devstore/internal/pkg/config"
	"github.com/ervinismu/devstore/internal/pkg/db"
	"github.com/ervinismu/devstore/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	cfg    config.Config
	DBConn *sqlx.DB
	err    error
)

func init() {

	// read configuration
	cfg, err = config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load app config")
	}

	// connect database
	DBConn, err = db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic("db not established")
	}

	// setup logrus
	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	r := gin.New()

	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// endpoints
	categoryRepository := repository.NewCategoryRepository(DBConn)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	r.GET("/categories", categoryController.BrowseCategory)
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories/:id", categoryController.DetailCategory)

	// update article by id
	r.PATCH("/categories/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "update article by id"})
	})

	// delete article by id
	r.DELETE("/categories/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "delete article by id"})
	})

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
