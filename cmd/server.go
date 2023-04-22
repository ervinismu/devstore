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
	db     *sqlx.DB
	router *gin.Engine
}

func NewServer(config config.Config, db *sqlx.DB) (*Server, error) {
	server := &Server{
		cfg: config,
		db:  db,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	// repo
	categoryRepository := repository.NewCategoryRepository(server.db)
	productRepository := repository.NewProductRepository(server.db)
	userRepository := repository.NewUserRepository(server.db)
	authRepository := repository.NewAuthRepository(server.db)

	// service
	tokenMaker := service.NewTokenMaker(
		server.cfg.AccessTokenKey,
		server.cfg.RefreshTokenKey,
		server.cfg.AccessTokenDuration,
		server.cfg.RefreshTokenDuration,
	)

	categoryService := service.NewCategoryService(categoryRepository)
	registrationService := service.NewRegistrationService(userRepository)
	sessionService := service.NewSessionService(userRepository, authRepository, tokenMaker)
	productService := service.NewProductService(productRepository, categoryRepository)

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

	authRouters := router.Group("/").Use(
		middleware.AuthMiddleware(tokenMaker),
	)

	authRouters.GET("/auth/logout", sessionController.Logout)
	authRouters.GET("/categories", categoryController.BrowseCategory)
	authRouters.POST("/categories", categoryController.CreateCategory)
	authRouters.GET("/categories/:id", categoryController.DetailCategory)
	authRouters.DELETE("/categories/:id", categoryController.DeleteCategory)
	authRouters.PATCH("/categories/:id", categoryController.UpdateCategory)
	authRouters.GET("/products", productController.BrowseProduct)
	authRouters.POST("/products", productController.CreateProduct)
	authRouters.GET("/products/:id", productController.DetailProduct)
	authRouters.DELETE("/products/:id", productController.DeleteProduct)
	authRouters.PATCH("/products/:id", productController.UpdateProduct)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
