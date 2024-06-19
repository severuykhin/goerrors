package goerrors

import (
	"encoding/json"
	"fmt"

	"github.com/severuykhin/gostacktrace"
)

type (
	kind string
	err  struct {
		kind        kind
		code        string
		message     string
		stack       string
		context     map[string]any
		wrapedError error
		innerErrors multiError
	}

	errSerializer struct {
		Kind    string         `json:"kind"`
		Code    string         `json:"code"`
		Message string         `json:"message"`
		Stack   string         `json:"stack"`
		Context map[string]any `json:"context"`
	}
)

const (
	ErrInvalidValue kind = "ErrInvalidValue"
	ErrNotFound     kind = "ErrNotFound"
	ErrInternal     kind = "ErrInternal"
	ErrAccessDenied kind = "ErrAccessDenied"
	ErrConflict     kind = "ErrConflict"
	ErrMulti        kind = "ErrMulti"
)

func (e err) Error() string {

	code := ""
	if e.code != "" {
		code = "(" + e.code + ")"
	}
	message := fmt.Sprintf("[%s]%s: %s", string(e.kind), code, e.message)
	if e.wrapedError != nil {
		message += fmt.Sprintf("; %s", e.wrapedError.Error())
	}

	if e.innerErrors != nil {
		message += e.innerErrors.Error()
	}

	return message
}

func (e err) Wrap(er error) error {
	e.wrapedError = er
	return e
}

func (e err) Unwrap() error {
	return e.wrapedError
}

func (e err) WithContext(keyvals ...any) err {

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

func (e err) WithCode(c string) err {
	e.code = c
	return e
}

func (e err) WithStack(depth int) err {
	st := gostacktrace.Get(uint(depth))
	frames := st.GetFrames()
	var stData string
	if len(frames) < depth {
		stData = frames.ToString()
	} else {
		stData = frames[:depth].ToString()
	}
	e.stack = stData
	return e
}

func (e err) GetKind() kind {
	return e.kind
}

func (e err) GetCode() string {
	return e.code
}

func (e err) GetMessage() string {
	return e.message
}

func (e err) GetContext() context {
	return e.context
}

func (e err) GetInnerErrors() []err {
	return e.innerErrors
}

func (e err) GetStackTrace() string {
	return e.stack
}

func (e err) MarshalJSON() ([]byte, error) {
	return json.Marshal(&errSerializer{
		Kind:    string(e.kind),
		Code:    e.code,
		Message: e.message,
		Stack:   e.GetStackTrace(),
		Context: e.context,
	})
}

func (e *err) UnmarshalJSON(data []byte) error {
	var errS errSerializer
	if err := json.Unmarshal(data, &errS); err != nil {
		return err
	}

	e.kind = kind(errS.Kind)
	e.code = errS.Code
	e.message = errS.Message
	e.context = errS.Context
	e.stack = errS.Stack
	return nil
}
