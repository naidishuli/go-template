package apierror

func Unauthorized(details map[string]any, errs ...error) *Error {
	var err error
	if len(errs) > 0 {
		err = errs[0]
	}

	return &Error{
		Status: 401,
		Response: Response{
			Message: "unauthorized",
			Details: details,
		},
		err: err,
	}
}

func BadRequest(err error, details any) *Error {
	return &Error{
		Status: 400,
		Response: Response{
			Message: "bad request",
			Details: details,
		},
	}
}

func BadRequestMalformed(err error) *Error {
	return &Error{
		Status: 400,
		Response: Response{
			Message: "bad request",
			Details: map[string]interface{}{
				"error": "malformed request",
			},
		},
	}
}

func InternalServerError(err error) *Error {
	return &Error{
		Status: 500,
		Response: Response{
			Message: "internal server error",
		},
	}
}

func NoPermission() *Error {
	return &Error{
		Status: 401,
		Response: Response{
			Message: "user has no permission",
		},
	}
}

func NotFound() *Error {
	return &Error{
		Status: 404,
		Response: Response{
			Message: "resource not found",
			Details: nil,
		},
	}
}
