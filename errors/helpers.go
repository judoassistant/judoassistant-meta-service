package errors

import (
	"github.com/pkg/errors"
)

func New(msg string, code int) error {
	return &codedError{
		wrappedErr: errors.Errorf(msg).(StackTracer),
		code:       code,
	}
}

func Wrap(err error, msg string) error {
	return &codedError{
		msg:        msg,
		wrappedErr: withStack(err),
		code:       Code(err),
	}
}

func WrapCode(err error, msg string, code int) error {
	return &codedError{
		msg:        msg,
		wrappedErr: withStack(err),
		code:       code,
	}
}

func Code(err error) int {
	if coder, ok := err.(Coder); ok {
		return coder.Code()
	}

	return CodeInternal
}

func withStack(err error) StackTracer {
	if stackTracer, ok := err.(StackTracer); ok {
		return stackTracer
	}

	return errors.WithStack(err).(StackTracer)
}
