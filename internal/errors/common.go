package errors

func NotFound(data interface{}) error {
	return &Error{
		Code:    ResourceNotFound,
		Message: "resource not found",
		Details: data,
		Trace:   trace(),
	}
}
