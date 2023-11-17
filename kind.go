package goerrors

type (
	kind string
)

const (
	ErrBadRequest kind = "ErrBadRequest"
	ErrInternal   kind = "ErrInternal"
)
