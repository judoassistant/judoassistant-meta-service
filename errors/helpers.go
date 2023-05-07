package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func New(msg string, code int) error {
	return &codedError{
		wrappedErr: errors.Errorf(msg).(StackTracer),
		code:       code,
	}
}

func Wrap(err error, msg string) error {
	result := &codedError{
		msg:        msg,
		wrappedErr: withStack(err),
		code:       http.StatusInternalServerError,
	}
	if coder, ok := err.(Coder); ok {
		result.code = coder.Code()
	}

	return result
}

func WrapCode(err error, msg string, code int) error {
	return &codedError{
		msg:        msg,
		wrappedErr: withStack(err),
		code:       code,
	}
}

func withStack(err error) StackTracer {
	if stackTracer, ok := err.(StackTracer); ok {
		return stackTracer
	}

	return errors.WithStack(err).(StackTracer)
}
