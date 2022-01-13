package errors

func NotFound(data interface{}) error {
	return &Error{
		Code:    NotFoundError,
		Message: "resource not found",
		Details: data,
		Trace:   trace(),
	}
}
