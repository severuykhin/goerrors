package goerrors

import (
	"encoding/json"
	"testing"
)

func provideInvalidValueError(message string, code string) error {
	return NewInvalidValueErr(message).WithCode(code).WithStack(2)
}

func TestGoError(t *testing.T) {

	err := provideInvalidValueError("some message", "Email.Invalid")

	domainErr := From(err)

	if domainErr.GetKind() != ErrInvalidValue {
		t.Fatalf("wrong kind of error")
	}

	if domainErr.GetCode() != "Email.Invalid" {
		t.Fatalf("wrong error code")
	}

	data, err := json.Marshal(domainErr)
	if err != nil {
		t.Fatalf("error when serializing to json")
	}

	domainErrorParsed, parseErr := FromJson(data)
	if parseErr != nil {
		t.Fatalf("error when deserializing from json")
	}

	if domainErrorParsed.GetKind() != ErrInvalidValue {
		t.Fatalf("wrong kind of error")
	}

	if domainErrorParsed.GetCode() != "Email.Invalid" {
		t.Fatalf("wrong error code")
	}

	if domainErrorParsed.GetStackTrace() == "" {
		t.Fatalf("stack trace must not be empty")
	}
}

func TestMultiError(t *testing.T) {

	mError := NewMultiError()
	if mError != nil {
		t.Fatalf("Multierror should be nil")
	}

	err1 := provideInvalidValueError("email must match the pattern", "Email.Invalid")
	err2 := provideInvalidValueError("phone number must match the pattern", "Phone.Invalid")
	mError = NewMultiError(err1, err2)

	if mError == nil {
		t.Fatalf("Multierror should not be nil")
	}
}
