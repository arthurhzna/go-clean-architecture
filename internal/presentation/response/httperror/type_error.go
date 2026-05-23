package httperror

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/arthurhzna/go-clean-architecture/internal/presentation/response/httperror/constant"
)

func newConflictError(err error) *ResponseError {
	return NewResponseError(err, http.StatusConflict, fmt.Sprintf(constant.ConflictErrorMessage, err.Error()))
}

func newUnauthorizedError(err error) *ResponseError {
	return NewResponseError(err, http.StatusUnauthorized, constant.UnauthorizedErrorMessage)
}

func newNotFoundError(err error) *ResponseError {
	return NewResponseError(err, http.StatusNotFound, fmt.Sprintf(constant.NotFoundErrorMessage, err.Error()))
}

func newInternalServerError(err error) *ResponseError {
	return NewResponseError(err, http.StatusInternalServerError, constant.InternalServerErrorMessage)
}

func NewTimeoutError() *ResponseError {
	return NewResponseError(
		errors.New(constant.RequestTimeoutErrorMessage),
		http.StatusRequestTimeout,
		constant.RequestTimeoutErrorMessage,
	)
}
