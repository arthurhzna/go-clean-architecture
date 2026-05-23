package bootstrap

import (
	"net/http"

	"github.com/arthurhzna/go-clean-architecture/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/arthurhzna/go-clean-architecture/internal/bootstrap/middleware"
)

type HttpServer struct {
	cfg    *config.Config
	server *http.Server
	logger *logger.Logger
}


func NewHTTPServer(cfg *config.Config, logger *logger.Logger) *HttpServer {
	appCfg := cfg.App
	gin.SetMode(appCfg.Environment)
	router := gin.New()
	router.ContextWithFallback = true
	router.HandleMethodNotAllowed = true

	RegisterMiddleware(router, cfg)

	return &HttpServer{
		cfg:    cfg,
		server: &http.Server{
			Addr:    appCfg.HTTPAddr,
			Handler: router,
		},
		logger: logger,
	}
}

func (s *HttpServer) Start() {
	s.logger.Log.Info("Running HTTP server on port:", s.cfg.HttpServer.Port)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Log.Fatal("Error while HTTP server listening:", err)
	}
	s.logger.Log.Info("HTTP server is not receiving new requests...")
}

func (s *HttpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.HttpServer.GracePeriod)*time.Second)
	defer cancel()

	s.logger.Log.Info("Attempting to shut down the HTTP server...")
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Log.Fatal("Error shutting down HTTP server:", err)
	}
	s.logger.Log.Info("HTTP server shut down gracefully")
}


func RegisterMiddleware(router *gin.Engine, cfg *config.AppConfig) {
	gzip.Gzip(gzip.BestSpeed),
	middleware.Logger(),
	middleware.ErrorHandler(),
	middleware.RequestTimeout(cfg.App.RequestTimeout),
	cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}),
	gin.Recovery(),
}