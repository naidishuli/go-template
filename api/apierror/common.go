package apierror

import "go-template/internal/app"

func Unauthorized(err error, details map[string]interface{}) *Error {
	return &Error{
		Status: 401,
		Response: Response{
			Message: "Unauthorized",
			Details: map[string]interface{}{
				"reason": details["reason"],
			},
		},
		err: app.NewError(app.ErrUndefined, err, details),
	}
}

func BadRequest(err error, details interface{}) *Error {
	return &Error{
		Status: 400,
		Response: Response{
			Message: "Bad request",
			Details: details,
		},
		err: app.NewError(app.ErrUndefined, err, nil),
	}
}

func BadRequestMalformed(err error) *Error {
	return &Error{
		Status: 400,
		Response: Response{
			Message: "Bad request",
			Details: map[string]interface{}{
				"error": "malformed request",
			},
		},
		err: app.NewError(app.ErrUndefined, err, nil),
	}
}

func InternalServerError(err error) *Error {
	return &Error{
		Status: 500,
		Response: Response{
			Message: "Internal server error",
		},
		err: app.NewError(app.ErrUndefined, err, nil),
	}
}

func NoPermission() *Error {
	return &Error{
		Status: 401,
		Response: Response{
			Message: "User has no permission",
		},
	}
}

func NotFound() *Error {
	return &Error{
		Status: 404,
		Response: Response{
			Message: "Resource not found",
			Details: nil,
		},
	}
}
