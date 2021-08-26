package entity

type ErrorBadRequest struct {
	Message error
}

func (e *ErrorBadRequest) Error() string {
	return e.Message.Error()
}

func NewErrorBadRequest(Message error) error {
	return &ErrorBadRequest{
		Message,
	}
}

type ErrorUnauthorized struct {
	Message error
}

func (e *ErrorUnauthorized) Error() string {
	return e.Message.Error()
}

func NewErrorUnauthorized(Message error) error {
	return &ErrorUnauthorized{
		Message,
	}
}

type ErrorForbidden struct {
	Message error
}

func (e *ErrorForbidden) Error() string {
	return e.Message.Error()
}

func NewErrorForbidden(Message error) error {
	return &ErrorForbidden{
		Message,
	}
}

type ErrorNotFound struct {
	Message error
}

func (e *ErrorNotFound) Error() string {
	return e.Message.Error()
}

func NewErrorNotFound(Message error) error {
	return &ErrorNotFound{
		Message,
	}
}

type ErrorConflict struct {
	Message error
}

func (e *ErrorConflict) Error() string {
	return e.Message.Error()
}

func NewErrorConflict(Message error) error {
	return &ErrorConflict{
		Message,
	}
}

type ErrorInternalServer struct {
	Message error
}

func (e *ErrorInternalServer) Error() string {
	return e.Message.Error()
}

func NewErrorInternalServer(Message error) error {
	return &ErrorInternalServer{
		Message,
	}
}
