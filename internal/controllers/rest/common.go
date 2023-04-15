package rest

import (
	"errors"
	"net/http"
)

type (
	Message map[string]string

	Response struct {
		Status  int            `json:"status"`
		Message Message        `json:"message"`
		Errors  []CaptureError `json:"errors,omitempty"`
		Data    interface{}    `json:"data,omitempty"`
		Meta    interface{}    `json:"meta,omitempty"`
	}

	CaptureError struct {
		MoreInfo        string `json:"moreInfo"`
		UserMessage     string `json:"userMessage"`
		InternalMessage string `json:"internalMessage"`
		Code            int    `json:"code"`
	}
)

//nolint:gochecknoglobals
var (
	text = http.StatusText

	msgSuccess    = map[string]string{"en": "Success", "id": "Sukses"}
	msgFailed     = map[string]string{"en": "Failed", "id": "Gagal"}
	msgUnexpected = map[string]string{
		"en": "An unexpected error occurred. Please try again later.",
		"id": "Terjadi kesalahan tak terduga. Silakan coba lagi nanti.",
	}
)

func unwrapFirstError(err error) string { return UnwrapAll(err).Error() }

// UnwrapAll will unwrap the underlying error until we get the first wrapped error.
func UnwrapAll(err error) error {
	for err != nil && errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	return err
}

func NewResponseError(statusCode int, message Message, userMessage, internalMessage, moreInfo string) Response {
	return Response{
		Status:  statusCode,
		Message: message,
		Errors: []CaptureError{
			{
				Code:            statusCode,
				UserMessage:     userMessage,
				InternalMessage: internalMessage,
				MoreInfo:        moreInfo,
			},
		},
	}
}
