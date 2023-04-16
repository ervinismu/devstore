package main

import (
	"fmt"
	"net/http"

	"github.com/ervinismu/devstore/internal/app/controller"
	"github.com/ervinismu/devstore/internal/app/repository"
	"github.com/ervinismu/devstore/internal/app/service"
	"github.com/ervinismu/devstore/internal/pkg/config"
	"github.com/ervinismu/devstore/internal/pkg/db"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/ervinismu/devstore/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
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
	r := gin.New()

	// implement middleware
	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	r.GET("/ping", func(ctx *gin.Context) {
		handler.ResponseSuccess(ctx, http.StatusOK, "pong", nil)
	})

	// repo
	categoryRepository := repository.NewCategoryRepository(DBConn)
	productRepository := repository.NewProductRepository(DBConn)
	userRepository := repository.NewUserRepository(DBConn)
	authRepository := repository.NewAuthRepository(DBConn)

	// service
	categoryService := service.NewCategoryService(categoryRepository)
	registrationService := service.NewRegistrationService(userRepository)
	sessionService := service.NewSessionService(userRepository, authRepository)
	productService := service.NewProductService(productRepository, categoryRepository)

	// controller
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	registrationController := controller.NewRegistrationController(registrationService)
	sessionController := controller.NewSessionController(sessionService)

	r.POST("/register", registrationController.Register)
	r.POST("/login", sessionController.SignIn)

	r.Use(middleware.AuthMiddleware())

	r.GET("/categories", categoryController.BrowseCategory)
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories/:id", categoryController.DetailCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)
	r.PATCH("/categories/:id", categoryController.UpdateCategory)

	r.GET("/products", productController.BrowseProduct)
	r.POST("/products", productController.CreateProduct)
	r.GET("/products/:id", productController.DetailProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
	r.PATCH("/products/:id", productController.UpdateProduct)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	err := r.Run(appPort)
	if err != nil {
		log.Panic(fmt.Errorf("error cannot start app : %w", err))
	}
}
