package errors

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type codedError struct {
	wrappedErr stackTracer
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

func (e *codedError) Format(s fmt.State, verb rune) {
	if verb == 'q' {
		fmt.Fprintf(s, "%q", e.Error())
		return
	}

	if verb == 'v' && s.Flag('+') {
		io.WriteString(s, e.Error())
		e.wrappedErr.StackTrace().Format(s, verb)
		return
	}

	io.WriteString(s, e.Error())
}

type coder interface {
	Code() int
}

type wrapper interface {
	Unwrap() error
}

type stackTracer interface {
	StackTrace() errors.StackTrace
	Error() string
}

var _ wrapper = (*codedError)(nil)
var _ coder = (*codedError)(nil)
var _ stackTracer = (*codedError)(nil)
var _ fmt.Formatter = (*codedError)(nil)
