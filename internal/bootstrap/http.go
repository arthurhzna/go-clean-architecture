package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/logger"

	"github.com/arthurhzna/go-clean-architecture/internal/config"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/controller"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	cfg    *config.Config
	server *http.Server
	logger logger.Logger
}

func NewHTTPServer(cfg *config.Config, logger logger.Logger) *HttpServer {
	appCfg := cfg.App
	httpCfg := cfg.HttpServer
	gin.SetMode(appCfg.Environment)
	router := gin.New()
	router.ContextWithFallback = true
	router.HandleMethodNotAllowed = true

	RegisterMiddleware(router, cfg.HttpServer, logger)

	return &HttpServer{
		cfg: cfg,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", httpCfg.Host, httpCfg.Port),
			Handler: router,
		},
		logger: logger,
	}
}

func (s *HttpServer) Start() {
	s.logger.Info("Running HTTP server on port:", s.cfg.HttpServer.Port)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Fatal("Error while HTTP server listening:", err)
	}
	s.logger.Info("HTTP server is not receiving new requests...")
}

func (s *HttpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.HttpServer.GracePeriod)*time.Second)
	defer cancel()

	s.logger.Info("Attempting to shut down the HTTP server...")
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Fatal("Error shutting down HTTP server:", err)
	}
	s.logger.Info("HTTP server shut down gracefully")
}

func RegisterMiddleware(router *gin.Engine, cfg *config.HttpServerConfig, logger logger.Logger) {
	middlewares := []gin.HandlerFunc{
		gzip.Gzip(gzip.BestSpeed),
		middleware.Logger(logger),
		middleware.ErrorHandler(),
		middleware.RequestTimeout(cfg.RequestTimeoutPeriod),
		cors.New(cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowAllOrigins:  true,
			AllowCredentials: true,
		}),
		gin.Recovery(),
	}

	router.Use(middlewares...)
}

func RegisterRoutes(
	router *gin.Engine,
	appController *controller.AppController,
) {
	router.NoRoute(appController.RouteNotFound)
	router.NoMethod(appController.MethodNotAllowed)

	api := router.Group("/apiv1")

	api.GET("/health", appController.Health)
}
