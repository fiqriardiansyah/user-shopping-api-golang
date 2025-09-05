package helper

import (
	"net/http"

	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/model"
	"github.com/sirupsen/logrus"
)

type AppError model.ErrorResponse

func (e *AppError) Error() string {
	return e.Message
}

func BadRequest(msg string) *AppError {
	logrus.Error("BAD REQUEST ", msg)
	return &AppError{
		Message: msg,
		Code:    http.StatusBadRequest,
		Status:  "BAD REQUEST",
	}
}

func Unauthorized(msg string) *AppError {
	logrus.Error("UNAUTHORIZED ", msg)
	return &AppError{
		Message: msg,
		Code:    http.StatusUnauthorized,
		Status:  "UNAUTHORIZED",
	}
}

func NotFound(msg string) *AppError {
	logrus.Error("NOT FOUND ", msg)
	return &AppError{
		Message: msg,
		Code:    http.StatusNotFound,
		Status:  "NOT FOUND",
	}
}

func Internal(msg string) *AppError {
	logrus.Error("INTERNAL SERVER ERROR ", msg)
	return &AppError{
		Message: msg,
		Code:    http.StatusInternalServerError,
		Status:  "INTERNAL SERVER ERROR",
	}
}

func Forbidden(msg string) *AppError {
	logrus.Error("FORBIDDEN ", msg)
	return &AppError{
		Message: msg,
		Code:    http.StatusForbidden,
		Status:  "FORBIDDEN",
	}
}
