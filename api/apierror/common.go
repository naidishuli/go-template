package apierror

func Unauthorized(details, devDetails interface{}) *ApiError {
	return &ApiError{
		Status:  401,
		Message: "Unauthorized",
		Details: details,
		Developer: &developer{
			Details: devDetails,
		},
	}
}

func BadRequest(details interface{}) *ApiError {
	return &ApiError{
		Status:  400,
		Message: "Bad request",
		Details: details,
	}
}
