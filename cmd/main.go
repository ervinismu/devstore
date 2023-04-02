package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ervinismu/devstore/internal/app/controller"
	"github.com/ervinismu/devstore/internal/app/repository"
	"github.com/ervinismu/devstore/internal/app/service"
	"github.com/ervinismu/devstore/internal/pkg/config"
	"github.com/ervinismu/devstore/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	categoryRepo := repository.NewCategoryRepository(DBConn)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	r.GET("/categories", categoryController.BrowseCategories)
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories/:id", categoryController.GetCategory)

	// update
	r.PATCH("/categories", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, gin.H { "message": "patch category" })
	})

	// delete
	r.DELETE("/categories/:id", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, gin.H { "message": "delete category" })
	})

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
