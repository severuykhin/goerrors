package goerrors

import (
	"testing"
)

func TestGoErrors(t *testing.T) {
	err := NewBadRequestErr()
	if err.kind != ErrBadRequest {
		t.Fatalf("wrong Kind of error")
	}

	strace := err.GetStack(2)
	if len(strace) == 0 {
		t.Fatalf("stacktrace is empty")
	}
}
