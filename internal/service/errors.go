package service

import "errors"

var (
	ErrUniqueViolation = errors.New("error already exists")

	ErrInvalidOffset = errors.New("invalid offset")
	ErrInvalidLimit  = errors.New("invalid limit")
)
