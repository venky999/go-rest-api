package server

import (
	"context"
	"go-rest-api/internal/handlers"
	zap_echo_logger "go-rest-api/internal/logger"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/validators"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
}

func New(db *gorm.DB, zapLogger *zap.Logger) *Server {

	e := echo.New()
	e.Validator = &validators.CustomValidator{Validator: validator.New()}
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "frame-ancestors 'none'; default-src 'none'",
	}))
	//e.Use(middleware.Secure())
	e.Use(zap_echo_logger.ZapLogger(zapLogger))

	txnHandler := handlers.NewTransactionHandler(repository.NewTransactionRepository(db, zapLogger))

	// Setup routes
	api := e.Group("/api")
	api.POST("/transaction", txnHandler.InsertTransaction)
	api.POST("/transaction/", txnHandler.InsertTransaction)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return &Server{
		echo: e,
	}
}

func (s *Server) Start() error {
	return s.echo.Start(":8080")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
