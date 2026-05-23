package error

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response/error/constant"
)

func NewNotFoundError(msg string) *ResponseError {
	msg = fmt.Sprintf(constant.NotFoundErrorMessage, msg)
	err := errors.New(msg)

	return NewResponseError(err, http.StatusNotFound, msg)
}
