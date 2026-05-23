package error

import (
	"errors"
	"net/http"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response/error/constant"
)

func NewTimeoutError() *ResponseError {
	msg := constant.RequestTimeoutErrorMessage

	err := errors.New(msg)

	return NewResponseError(err, http.StatusRequestTimeout, msg)
}
