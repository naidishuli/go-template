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

func BadRequest(err error, details interface{}) *ApiError {
	apiError := &ApiError{
		Status:  400,
		Message: "BadRequest",
		Details: details,
	}

	if err != nil {
		apiError.Developer = &developer{
			Details: err,
		}
	}

	return apiError
}
