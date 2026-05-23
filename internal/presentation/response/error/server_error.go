package error

import (
	"errors"
	"net/http"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response/error/constant"
)

func NewServerError() *ResponseError {
	msg := constant.InternalServerErrorMessage

	err := errors.New(msg)

	return NewResponseError(err, http.StatusInternalServerError, msg)
}
