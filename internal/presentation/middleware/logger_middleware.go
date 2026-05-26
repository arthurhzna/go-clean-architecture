package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/arthurhzna/go-clean-architecture/internal/domain/logger"
	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response"
)

func Logger(logger logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		ctx.Next()

		params := map[string]any{
			"status_code": ctx.Writer.Status(),
			"client_ip":   ctx.ClientIP(),
			"method":      ctx.Request.Method,
			"latency":     time.Since(start).String(),
			"path":        path,
		}

		if len(ctx.Errors) == 0 {
			logger.WithFields(params).Info("incoming request")
			return
		}
		logErrors(ctx, params, logger)
	}
}

// todo custom validate
func logErrors(ctx *gin.Context, params map[string]any, logger logger.Logger) {
	errors := []error{}
	for _, err := range ctx.Errors {
		switch e := err.Err.(type) {
		case validator.ValidationErrors:
			params["status_code"] = http.StatusBadRequest
			errors = append(errors, err)
		case *response.ResponseError:
			params["status_code"] = e.GetCode()
			errors = append(errors, e.OriginalError())
		default:
			params["status_code"] = http.StatusInternalServerError
			errors = append(errors, err)
		}
	}

	params["errors"] = errors
	logger.WithFields(params).Error("got error")
}
