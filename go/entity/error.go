package entity

type ErrorBadRequest struct {
	message error
}

func (e *ErrorBadRequest) Error() string {
	return e.message.Error()
}

func NewErrorBadRequest(message error) error {
	return &ErrorBadRequest{
		message,
	}
}

type ErrorUnauthorized struct {
	message error
}

func (e *ErrorUnauthorized) Error() string {
	return e.message.Error()
}

func NewErrorUnauthorized(message error) error {
	return &ErrorUnauthorized{
		message,
	}
}

type ErrorForbidden struct {
	message error
}

func (e *ErrorForbidden) Error() string {
	return e.message.Error()
}

func NewErrorForbidden(message error) error {
	return &ErrorForbidden{
		message,
	}
}

type ErrorNotFound struct {
	message error
}

func (e *ErrorNotFound) Error() string {
	return e.message.Error()
}

func NewErrorNotFound(message error) error {
	return &ErrorNotFound{
		message,
	}
}

type ErrorConflict struct {
	message error
}

func (e *ErrorConflict) Error() string {
	return e.message.Error()
}

func NewErrorConflict(message error) error {
	return &ErrorConflict{
		message,
	}
}

type ErrorInternalServer struct {
	message error
}

func (e *ErrorInternalServer) Error() string {
	return e.message.Error()
}

func NewErrorInternalServer(message error) error {
	return &ErrorInternalServer{
		message,
	}
}
