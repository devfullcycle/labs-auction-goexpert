package internal_error

type InternalError struct {
	Message string
	Err     string
}

func (ie *InternalError) Error() string {
	return ie.Message
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not_found",
	}
}

func NewInternalServerError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "internal_server_error",
	}
}

func NewBadRequestError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "bad_request",
	}
}
