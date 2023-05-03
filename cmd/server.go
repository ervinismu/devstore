package main

import (
	"net/http"

	"github.com/ervinismu/devstore/internal/app/controller"
	"github.com/ervinismu/devstore/internal/app/repository"
	"github.com/ervinismu/devstore/internal/app/service"
	"github.com/ervinismu/devstore/internal/pkg/config"
	"github.com/ervinismu/devstore/internal/pkg/handler"
	"github.com/ervinismu/devstore/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	cfg    config.Config
	dbConn *sqlx.DB
	router *gin.Engine
}

func NewServer(cfg config.Config, dbConn *sqlx.DB) (*Server, error) {
	server := &Server{
		cfg:    cfg,
		dbConn: dbConn,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	// repo
	categoryRepository := repository.NewCategoryRepository(server.dbConn)
	productRepository := repository.NewProductRepository(server.dbConn)
	userRepository := repository.NewUserRepository(server.dbConn)
	authRepository := repository.NewAuthRepository(server.dbConn)

	// service
	tokenMaker := service.NewTokenMaker(
		server.cfg.AccessTokenKey,
		server.cfg.RefreshTokenKey,
		server.cfg.AccessTokenDuration,
		server.cfg.RefreshTokenDuration,
	)
	uploaderService := service.NewUploaderService(
		server.cfg.CloudinaryCloudName,
		server.cfg.CloudinaryApiKey,
		server.cfg.CloudinaryApiSecret,
		server.cfg.CloudinaryUploadFolder,
	)
	categoryService := service.NewCategoryService(categoryRepository)
	registrationService := service.NewRegistrationService(userRepository)
	productService := service.NewProductService(productRepository, categoryRepository, uploaderService)
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

	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", registrationController.Register)
	authRoutes.POST("/login", sessionController.Login)
	authRoutes.GET("/refresh", sessionController.Refresh)
	authRoutes.GET("/logout", middleware.AuthMiddleware(tokenMaker), sessionController.Logout)

	router.Use(middleware.AuthMiddleware(tokenMaker))

	router.GET("/products", productController.BrowseProduct)
	router.POST("/products", productController.CreateProduct)
	router.GET("/products/:id", productController.DetailProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
	router.PATCH("/products/:id", productController.UpdateProduct)

	router.GET("/categories", categoryController.BrowseCategory)
	router.POST("/categories", categoryController.CreateCategory)
	router.GET("/categories/:id", categoryController.DetailCategory)
	router.DELETE("/categories/:id", categoryController.DeleteCategory)
	router.PATCH("/categories/:id", categoryController.UpdateCategory)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
