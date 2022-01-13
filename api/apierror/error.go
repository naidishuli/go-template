package apierror

import (
	"go-template/internal/errors"
)

// ApiError represent an errors to be returned to the client
type ApiError struct {
	Status    int         `json:"-"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Developer *developer  `json:"developer,omitempty"`
}

// developer contains information for development purpose only, this info should not pass to client
type developer struct {
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

// Error used to implement errors interface
func (e *ApiError) Error() string {
	return e.Message
}

// New return a new ApiError based on the errors passed on argument
func New(err error) *ApiError {
	var apiError *ApiError

	switch err.(type) {
	case *errors.Error:
		apiError = apiErrorFromInternalError(err.(*errors.Error))
	default:
		apiError = &ApiError{
			Status:  500,
			Message: "Internal server errors",
		}
	}

	return apiError
}

// apiErrorFromInternalError transform internal error to an api error
func apiErrorFromInternalError(err *errors.Error) *ApiError {
	if err == nil {
		return &ApiError{
			Status:  500,
			Message: "Internal server error",
		}
	}

	apiError := &ApiError{
		Message: err.Message,
		Details: err.Details,
	}

	if err.Developer != nil {
		apiError.Developer = &developer{
			Message: err.Developer.Message,
			Details: err.Developer.Details,
		}
	}

	switch err.Code {
	case errors.GenerateToken:
		apiError.Status = 500
		apiError.Message = "Internal server error"
	case errors.ResourceNotFound:
		apiError.Status = 404
		apiError.Message = "Resource not found"
	default:
		apiError.Status = 500
		apiError.Message = "Internal server error"
	}

	return &ApiError{}
}
