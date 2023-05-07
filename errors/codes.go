package errors

import (
	"net/http"
)

const (
	CodeOK = http.StatusOK

	CodeBadRequest   = http.StatusBadRequest
	CodeUnauthorized = http.StatusUnauthorized
	CodeForbidden    = http.StatusForbidden
	CodeNotFound     = http.StatusNotFound

	CodeInternal       = http.StatusInternalServerError
	CodeNotImplemented = http.StatusNotImplemented
	CodeUnavailable    = http.StatusServiceUnavailable
)
