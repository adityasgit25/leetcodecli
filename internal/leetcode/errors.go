package leetcode

import (
	"errors"
	"fmt"
)

type ErrorKind string

const (
	ErrorKindNotFound          ErrorKind = "not_found"
	ErrorKindUnavailable       ErrorKind = "unavailable"
	ErrorKindEndpointFailure   ErrorKind = "endpoint_failure"
	ErrorKindRateLimited       ErrorKind = "rate_limited"
	ErrorKindMalformedResponse ErrorKind = "malformed_response"
	ErrorKindMissingStats      ErrorKind = "missing_stats"
)

func (kind ErrorKind) String() string {
	return string(kind)
}

type Error struct {
	Kind ErrorKind
	Err  error
}

func (err *Error) Error() string {
	if err.Err == nil {
		return string(err.Kind)
	}
	return fmt.Sprintf("%s: %v", err.Kind, err.Err)
}

func (err *Error) Unwrap() error {
	return err.Err
}

func IsErrorKind(err error, kind ErrorKind) bool {
	var leetcodeErr *Error
	return errors.As(err, &leetcodeErr) && leetcodeErr.Kind == kind
}

func classify(kind ErrorKind, err error) error {
	return &Error{Kind: kind, Err: err}
}
