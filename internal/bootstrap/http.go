package bootstrap

import (
	"net/http"

	"github.com/arthurhzna/go-clean-architecture/internal/config"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(cfg *config.Config) *http.Server {
	appCfg := cfg.App
	gin.SetMode(appCfg.Environment)
	router := gin.New()
	router.ContextWithFallback = true
	router.HandleMethodNotAllowed = true
}
