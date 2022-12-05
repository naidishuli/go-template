package apierror

import (
	"encoding/json"
	"fmt"
	"go-template/internal/app"
	"go-template/internal/config"
)

type Response struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error represent an errors to be returned to the client
type Error struct {
	Status int `json:"-"`
	Response

	err *app.Error
}

// Error used to implement errors interface
func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Log() string {
	var code, message, trace, sql string

	if e.err != nil {
		code = string(e.err.Code)
		message = e.err.Message
		trace = e.err.Trace
		sql, _ = e.err.ErrData[config.GormSQLKey].(string)
	} else {
		data, _ := json.Marshal(e.Response)
		message = string(data)
	}

	return fmt.Sprintf(
		"Code: %s\n\nMessage: %s\n\nStack: \t%s\n\nSql: \t%s\n"+
			"--------------------------------------------------------------------------------------------------------------\n",
		code,
		message,
		trace,
		sql,
	)
}

// New return a new ApiError based on the errors passed on argument
func New(err error) *Error {
	var apiError *Error

	switch err.(type) {
	case *app.Error:
		apiError = apiErrorFromError(err.(*app.Error))
	default:
		apiError = &Error{
			Status: 500,
			Response: Response{
				Message: "Internal server error",
			},
			err: app.NewError(app.UndefinedErr, err, nil),
		}
	}

	return apiError
}

// apiErrorFromError transform internal error to an api error
func apiErrorFromError(err *app.Error) *Error {
	if err == nil {
		return nil
	}

	apiError := &Error{
		err: err,
	}

	switch err.Code {
	case app.NotFoundErr:
		apiError.Status = 404
		apiError.Message = err.Message
	default:
		apiError.Status = 500
		apiError.Message = "Internal server error"
	}

	return apiError
}
