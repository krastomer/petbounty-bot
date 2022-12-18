package command

import "errors"

var (
	ErrBadRequest           = errors.New("bad request")
	ErrInternalServiceError = errors.New("internal service error")
)
