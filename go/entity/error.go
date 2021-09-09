package entity

type ErrorBadRequest struct {
	Message error
	Code    string
}

func (e *ErrorBadRequest) Error() string {
	return e.Message.Error()
}

func NewErrorBadRequest(message error, code string) error {
	return &ErrorBadRequest{
		Message: message,
		Code:    code,
	}
}

type ErrorUnauthorized struct {
	Message error
	Code    string
}

func (e *ErrorUnauthorized) Error() string {
	return e.Message.Error()
}

func NewErrorUnauthorized(message error, code string) error {
	return &ErrorUnauthorized{
		Message: message,
		Code:    code,
	}
}

type ErrorForbidden struct {
	Message error
	Code    string
}

func (e *ErrorForbidden) Error() string {
	return e.Message.Error()
}

func NewErrorForbidden(message error, code string) error {
	return &ErrorForbidden{
		Message: message,
		Code:    code,
	}
}

type ErrorNotFound struct {
	Message error
	Code    string
}

func (e *ErrorNotFound) Error() string {
	return e.Message.Error()
}

func NewErrorNotFound(message error, code string) error {
	return &ErrorNotFound{
		Message: message,
		Code:    code,
	}
}

type ErrorConflict struct {
	Message error
	Code    string
}

func (e *ErrorConflict) Error() string {
	return e.Message.Error()
}

func NewErrorConflict(message error, code string) error {
	return &ErrorConflict{
		Message: message,
		Code:    code,
	}
}

type ErrorInternalServer struct {
	Message error
	Code    string
}

func (e *ErrorInternalServer) Error() string {
	return e.Message.Error()
}

func NewErrorInternalServer(message error, code string) error {
	return &ErrorInternalServer{
		Message: message,
		Code:    code,
	}
}
