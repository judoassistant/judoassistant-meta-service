package errors

import (
	"net/http"
)

const (
	CodeOK = http.StatusOK

	CodeBadRequest   = http.StatusBadRequest
	CodeConflict     = http.StatusConflict
	CodeUnauthorized = http.StatusUnauthorized
	CodeForbidden    = http.StatusForbidden
	CodeNotFound     = http.StatusNotFound

	CodeInternal       = http.StatusInternalServerError
	CodeNotImplemented = http.StatusNotImplemented
	CodeUnavailable    = http.StatusServiceUnavailable
)

func IsServerSide(code int) bool {
	return code >= 500
}
