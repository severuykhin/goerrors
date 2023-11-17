package goerrors

import (
	"fmt"

	"github.com/severuykhin/gostacktrace"
)

type (
	err struct {
		kind    kind
		message string
		stack   gostacktrace.StackTrace
	}
)

func (e err) Error() string {
	return fmt.Sprintf("[%s]: %s", e.kind, e.message)
}

func (e err) WithMessage(message string) err {
	e.message = message
	return e
}

func (e err) Produce(code string) {

}

func (e err) GetKind() kind {
	return e.kind
}

func (e err) GetMessage() string {
	return e.message
}

func (e err) GetStack(depth int) string {
	frames := e.stack.GetFrames()
	if len(frames) < depth {
		return frames.ToString()
	}
	return frames[:depth].ToString()
}

func From(er error) err {
	switch e := er.(type) {
	case err:
		return e
	default:
		return NewInternalErr().WithMessage(er.Error())
	}
}

func Is(err error, targetErrKind kind) bool {
	appErr := From(err)
	return appErr.kind == targetErrKind
}

func NewBadRequestErr() err {
	return err{
		kind:  ErrBadRequest,
		stack: gostacktrace.Get(3),
	}
}

func NewInternalErr() err {
	return err{
		kind:  ErrBadRequest,
		stack: gostacktrace.Get(3),
	}
}
