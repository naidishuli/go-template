package apierror

func Unauthorized(details interface{}) *ApiError {
	return &ApiError{
		Status:  401,
		Message: "Unauthorized",
		Details: details,
	}
}

func BadRequest(details interface{}) *ApiError {
	return &ApiError{
		Status:  400,
		Message: "Bad request",
		Details: details,
	}
}
