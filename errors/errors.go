package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type Coder interface {
	Code() int
}

type Wrapper interface {
	Unwrap() error
}

type StackTracer interface {
	StackTrace() errors.StackTrace
	Error() string
}

type codedError struct {
	wrappedErr StackTracer
	code       int
	msg        string
}

func (e *codedError) Code() int {
	return e.code
}

func (e *codedError) Error() string {
	if e.msg == "" {
		return e.wrappedErr.Error()
	}
	return fmt.Sprintf("%s: %s", e.msg, e.wrappedErr)
}

func (e *codedError) Unwrap() error {
	return e.wrappedErr
}

func (e *codedError) StackTrace() errors.StackTrace {
	return e.wrappedErr.StackTrace()
}

var _ Wrapper = (*codedError)(nil)
var _ Coder = (*codedError)(nil)
var _ error = (*codedError)(nil)
var _ StackTracer = (*codedError)(nil)
