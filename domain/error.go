package domain

import "errors"

var (
	// ErrInternalServerError -
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound -
	ErrNotFound = errors.New("your requested item is not found")
	// ErrConflict -
	ErrConflict = errors.New("your item already exist")
	// ErrBadParamInput -
	ErrBadParamInput = errors.New("given param is not valid")
)
