package errors

import (
	"github.com/pkg/errors"
)

func New(msg string, code int) error {
	return &codedError{
		wrappedErr: toStackTracer(errors.Errorf(msg)),
		code:       code,
	}
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	return &codedError{
		msg:        msg,
		wrappedErr: toStackTracer(err),
		code:       Code(err),
	}
}

func WrapWithCode(err error, msg string, code int) error {
	if err == nil {
		return nil
	}

	return &codedError{
		msg:        msg,
		wrappedErr: toStackTracer(err),
		code:       code,
	}
}

func Code(err error) int {
	if err == nil {
		return CodeOK
	}

	if coder, ok := err.(coder); ok {
		return coder.Code()
	}
	return CodeInternal
}

func toStackTracer(err error) stackTracer {
	if stackTracer, ok := err.(stackTracer); ok {
		return stackTracer
	}

	return errors.WithStack(err).(stackTracer)
}
