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

func NewServer(cfg config.Config, DBConn *sqlx.DB) (*Server, error) {
	server := &Server{
		cfg:    cfg,
		dbConn: DBConn,
	}

	// setup router
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	// repo
	categoryRepository := repository.NewCategoryRepository(server.dbConn)
	productRepository := repository.NewProductRepository(server.dbConn)
	userRepository := repository.NewUserRepository(server.dbConn)
	authRepository := repository.NewAuthRepository(server.dbConn)
	cartRepository := repository.NewCartRepository(server.dbConn)
	cartItemRepository := repository.NewCartItemRepository(server.dbConn)
	orderRepository := repository.NewOrderRepository(server.dbConn)

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
	midtransService := service.NewMidtransService(
		server.cfg.MidtransServerKey,
		server.cfg.MidtransMerchantID,
		server.cfg.MidtransBaseURL,
	)
	orderService := service.NewOrderService(orderRepository, midtransService)
	categoryService := service.NewCategoryService(categoryRepository)
	registrationService := service.NewRegistrationService(userRepository)
	productService := service.NewProductService(productRepository, categoryRepository, uploaderService)
	sessionService := service.NewSessionService(userRepository, authRepository, tokenMaker)
	cartService := service.NewCartService(productRepository, cartRepository, cartItemRepository)

	// controller
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	registrationController := controller.NewRegistrationController(registrationService)
	sessionController := controller.NewSessionController(sessionService, tokenMaker)
	cartController := controller.NewCartController(cartService)
	orderController := controller.NewOrderController(orderService)

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

	router.GET("/categories",
		middleware.PaginationMiddleware(cfg.PaginateDefaultPage, cfg.PaginateDefaultPageSize),
		categoryController.BrowseCategory)
	router.POST("/categories", categoryController.CreateCategory)
	router.GET("/categories/:id", categoryController.DetailCategory)
	router.DELETE("/categories/:id", categoryController.DeleteCategory)
	router.PATCH("/categories/:id", categoryController.UpdateCategory)

	router.GET("/products",
		middleware.PaginationMiddleware(cfg.PaginateDefaultPage, cfg.PaginateDefaultPageSize),
		productController.BrowseProduct)
	router.POST("/products", productController.CreateProduct)
	router.GET("/products/:id", productController.DetailProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)
	router.PATCH("/products/:id", productController.UpdateProduct)

	router.POST("/carts", cartController.AddToCart)

	router.GET("/orders/checkout", orderController.Checkout)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
