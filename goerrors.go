package goerrors

import (
	"fmt"
	"strconv"

	"github.com/severuykhin/gostacktrace"
)

type (
	err struct {
		kind    kind
		message string
		stack   gostacktrace.StackTrace
		context map[string]interface{}
	}
)

func (e err) Error() string {
	return fmt.Sprintf("[%s]: %s", e.kind, e.message)
}

func (e err) WithMessage(message string) err {
	e.message = message
	return e
}

func (e err) WithContext(keyvals ...interface{}) err {

	keyvalsLength := len(keyvals)

	if keyvalsLength == 0 {
		return e
	}

	errCtx := context{}

	for i := 0; i < keyvalsLength; i++ {
		if i > 0 && i%2 != 0 {
			continue
		}

		if i == keyvalsLength-1 {
			errCtx[valueToString(keyvals[i])] = ""
		} else {
			errCtx[valueToString(keyvals[i])] = valueToString(keyvals[i+1])
		}
	}

	e.context = errCtx

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

func (e err) GetContext() context {
	return e.context
}

func (e err) GetStackTrace(depth int) string {
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
		kind:  ErrInternal,
		stack: gostacktrace.Get(3),
	}
}

func NewNotFoundErr() err {
	return err{
		kind:  ErrNotFound,
		stack: gostacktrace.Get(3),
	}
}

func valueToString(val interface{}) string {
	switch v := val.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case error:
		return v.Error()
	default:
		return "unknowntype"
	}
}
