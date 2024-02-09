package response

import (
	"errors"
	"github.com/constantincuy/knowledgestore/internal/core/service"
	"net/http"
)

type ErrorDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func FromError(err error) Response {
	errorDTO := ErrorDTO{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}

	var serviceErr service.Error
	if errors.As(err, &serviceErr) {
		errorDTO.Message = serviceErr.CausedBy().Error()
		switch serviceErr.ServiceError() {
		case service.ErrBadRequest:
			errorDTO.Status = http.StatusBadRequest
		case service.ErrConflict:
			errorDTO.Status = http.StatusConflict
		case service.ErrForbidden:
			errorDTO.Status = http.StatusForbidden
		case service.ErrUnauthorized:
			errorDTO.Status = http.StatusUnauthorized
		case service.ErrInternalFailure:
			errorDTO.Status = http.StatusInternalServerError
		case service.ErrNotFound:
			errorDTO.Status = http.StatusNotFound
		}
	}

	return New().Status(errorDTO.Status).Json(errorDTO)
}

func FromErrorWithStatus(err error, status int) Response {
	errorDTO := ErrorDTO{
		Status:  status,
		Message: err.Error(),
	}
	return New().Status(errorDTO.Status).Json(errorDTO)
}
