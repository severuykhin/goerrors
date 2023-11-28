package goerrors

import (
	"testing"
)

func TestGoErrors(t *testing.T) {
	err := NewBadRequestErr().WithMessage("some message").WithContext("key1", "value1", "key2", 123, "key3", 444, 111, 222)
	if err.kind != ErrBadRequest {
		t.Fatalf("wrong kind of error")
	}

	strace := err.GetStackTrace(2)
	if len(strace) == 0 {
		t.Fatalf("stacktrace is empty")
	}

	errCtx := err.GetContext()
	errCtxList := errCtx.ToList()
	if len(errCtxList) != 8 {
		t.Fatalf("wrong number of context keyvalues")
	}
}
