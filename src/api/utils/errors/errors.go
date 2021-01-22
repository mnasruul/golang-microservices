package errors

import (
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	status  int    `json:"status"`
	message string `json:"message"`
	error   string `json:"error"`
}

func (e *apiError) Status() int {
	return e.Status()
}

func (e *apiError) Message() string {
	return e.Message()
}

func (e *apiError) Error() string {
	return e.Error()
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		status:  statusCode,
		message: message,
	}
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		status:  http.StatusNotFound,
		message: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		status:  http.StatusInternalServerError,
		message: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		status:  http.StatusBadRequest,
		message: message,
	}
}
