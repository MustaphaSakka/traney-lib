package exception

import "net/http"

type AppException struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func NotFoundException(message string) *AppException {
	return &AppException{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func InternalServerException(message string) *AppException {
	return &AppException{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func ValidationError(message string) *AppException {
	return &AppException{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func AuthenticationError(message string) *AppException {
	return &AppException{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func AuthorizationError(message string) *AppException {
	return &AppException{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

//Send only the exception message
//Code is already coming in the HTTP response
func (e AppException) AsMessage() *AppException {
	return &AppException{
		Message: e.Message,
	}
}
