package entity

type ErrorBadRequest struct {
	message string
}

func (e *ErrorBadRequest) Error() string {
	return e.message
}

func NewErrorBadRequest(message string) error {
	return &ErrorBadRequest{
		message: message,
	}
}

type ErrorUnauthorized struct {
	message string
}

func (e *ErrorUnauthorized) Error() string {
	return e.message
}

func NewErrorUnauthorized(message string) error {
	return &ErrorUnauthorized{
		message: message,
	}
}

type ErrorForbidden struct {
	message string
}

func (e *ErrorForbidden) Error() string {
	return e.message
}

func NewErrorForbidden(message string) error {
	return &ErrorForbidden{
		message: message,
	}
}

type ErrorNotFound struct {
	message string
}

func (e *ErrorNotFound) Error() string {
	return e.message
}

func NewErrorNotFound(message string) error {
	return &ErrorNotFound{
		message: message,
	}
}

type ErrorInternalServer struct {
	message string
}

func (e *ErrorInternalServer) Error() string {
	return e.message
}

func NewErrorInternalServer(message string) error {
	return &ErrorInternalServer{
		message: message,
	}
}
