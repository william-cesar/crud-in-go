package ierrors

import "net/http"

var errorTypes = map[string]string{
	"badRequest":    "We could not process your request. Please verify your fields.",
	"unauthorized":  "You must be logged in to perform this action.",
	"forbidden":     "You do not have permission to perform this action.",
	"notFound":      "Nothing was found with the given parameters.",
	"internalError": "The server could not process your request. Please try again later.",
}

type TError struct {
	Message    string   `json:"message"`
	StatusCode int      `json:"statusCode"`
	Causes     []TCause `json:"causes,omitempty"`
}

type TCause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *TError) Error() string {
	return e.Message
}

func NewError(message string, statusCode int) *TError {
	return &TError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewErrorWithCauses(causes []TCause) *TError {
	message := errorTypes["badRequest"]

	return &TError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Causes:     causes,
	}
}

func NewBadRequestError(message ...string) *TError {
	msg := errorTypes["badRequest"]

	if len(message) > 0 {
		msg = message[0]
	}

	return NewError(msg, http.StatusBadRequest)
}

func NewUnauthorizedError(message ...string) *TError {
	msg := errorTypes["unauthorized"]

	if len(message) > 0 {
		msg = message[0]
	}

	return NewError(msg, http.StatusUnauthorized)
}

func NewForbiddenError(message ...string) *TError {
	msg := errorTypes["forbidden"]

	if len(message) > 0 {
		msg = message[0]
	}

	return NewError(msg, http.StatusForbidden)
}

func NewNotFoundError(message ...string) *TError {
	msg := errorTypes["notFound"]

	if len(message) > 0 {
		msg = message[0]
	}

	return NewError(msg, http.StatusNotFound)
}

func NewInternalError(message ...string) *TError {
	msg := errorTypes["internalError"]

	if len(message) > 0 {
		msg = message[0]
	}

	return NewError(msg, http.StatusInternalServerError)
}
