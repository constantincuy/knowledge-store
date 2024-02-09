package service

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrNotFound        = errors.New("not found")
	ErrConflict        = errors.New("conflict")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrInternalFailure = errors.New("internal server error")
)

type Error struct {
	causedBy     error
	serviceError error
}

func (e Error) CausedBy() error {
	return e.causedBy
}

func (e Error) ServiceError() error {
	return e.serviceError
}

func (e Error) Error() string {
	return errors.Join(e.serviceError, e.causedBy).Error()
}

func NewError(serviceError, causedBy error) error {
	return Error{
		causedBy:     causedBy,
		serviceError: serviceError,
	}
}
