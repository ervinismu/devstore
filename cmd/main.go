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
	// repo
	categoryRepository := repository.NewCategoryRepository(DBConn)
	productRepository := repository.NewProductRepository(DBConn)
	userRepository := repository.NewUserRepository(DBConn)
	authRepository := repository.NewAuthRepository(DBConn)

	// service
	tokenMaker := service.NewTokenMaker(
		cfg.AccessTokenKey,
		cfg.RefreshTokenKey,
		cfg.AccessTokenDuration,
		cfg.RefreshTokenDuration,
	)
	categoryService := service.NewCategoryService(categoryRepository)
	registrationService := service.NewRegistrationService(userRepository)
	productService := service.NewProductService(productRepository, categoryRepository)
	sessionService := service.NewSessionService(userRepository, authRepository, tokenMaker)

	// controller
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	registrationController := controller.NewRegistrationController(registrationService)
	sessionController := controller.NewSessionController(sessionService, tokenMaker)

	router := gin.New()

	// implement middleware
	router.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	router.GET("/ping", func(ctx *gin.Context) {
		handler.ResponseSuccess(ctx, http.StatusOK, "pong", nil)
	})

	router.POST("/auth/register", registrationController.Register)
	router.POST("/auth/login", sessionController.Login)

	router.GET("/auth/refresh", sessionController.Refresh)

	// auth middleware
	router.Use(middleware.AuthMiddleware(tokenMaker))

	router.GET("/auth/logout", sessionController.Logout)
	router.GET("/categories", categoryController.BrowseCategory)
	router.POST("/categories", categoryController.CreateCategory)
	router.GET("/categories/:id", categoryController.DetailCategory)
	router.DELETE("/categories/:id", categoryController.DeleteCategory)
	router.PATCH("/categories/:id", categoryController.UpdateCategory)
	router.GET("/products", productController.BrowseProduct)
	router.POST("/products", productController.CreateProduct)
	router.GET("/products/:id", productController.DetailProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
	router.PATCH("/products/:id", productController.UpdateProduct)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	err := router.Run(appPort)
	if err != nil {
		log.Panic(fmt.Errorf("error cannot start app : %w", err))
	}
}
