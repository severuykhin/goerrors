package goerrors

type (
	kind string
)

const (
	ErrBadRequest   kind = "ErrBadRequest"
	ErrNotFound     kind = "ErrNotFound"
	ErrInternal     kind = "ErrInternal"
	ErrAccessDenied kind = "ErrAccessDenied"
	ErrConflict     kind = "ErrConflict"
)
