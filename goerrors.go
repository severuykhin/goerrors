package goerrors

import (
	"encoding/json"
)

func From(er error) err {
	switch e := er.(type) {
	case err:
		return e
	default:
		return NewInternalErr(er.Error())
	}
}

func FromJson(data []byte) (err, error) {
	e := err{}
	unmarshalError := json.Unmarshal(data, &e)
	return e, unmarshalError
}

func Is(err error, targetErrKind kind) bool {
	appErr := From(err)
	return appErr.kind == targetErrKind
}

func NewInvalidValueErr(message string) err {
	return err{
		kind:    ErrInvalidValue,
		message: message,
	}
}

func NewInternalErr(message string) err {
	return err{
		kind:    ErrInternal,
		message: message,
	}
}

func NewNotFoundErr(message string) err {
	return err{
		kind:    ErrNotFound,
		message: message,
	}
}

func NewAccessDeniedErr(message string) err {
	return err{
		kind:    ErrAccessDenied,
		message: message,
	}
}

func NewConflictErr(message string) err {
	return err{
		kind:    ErrConflict,
		message: message,
	}
}
