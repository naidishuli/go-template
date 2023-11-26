package apierror

import (
	"encoding/json"
	"errors"
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

	err error
}

// New return a new ApiError based on the errors passed on argument
func New(err error) *Error {
	var message string
	wErr := err

	for {
		appErr, ok := wErr.(*app.Error)
		if ok {
			return apiErrorFromAppError(appErr, message)
		}

		message = err.Error()
		wErr = errors.Unwrap(wErr)
		if wErr == nil {
			break
		}
	}

	dErr, ok := apiErrorFromDomainError(err)
	if ok {
		return dErr
	}

	return &Error{
		Status: 500,
		Response: Response{
			Message: "Internal server error",
		},
		err: &app.Error{Code: app.ErrUndefined, Err: err},
	}
}

// Error used to implement errors interface
func (e *Error) Error() string {
	return e.Message
}

// trace used to return the line of the code that creates the errors
// func trace() string {
//	pcs := make([]uintptr, 100)
//	runtime.Callers(2, pcs)
//	callersFrames := runtime.CallersFrames(pcs)
//
//	var trace string
//	for {
//		cf, more := callersFrames.Next()
//		if !more {
//			break
//		}
//
//		frameInfo := cf.Function
//		frameInfo = fmt.Sprintf("%s\n\t%s:%d", frameInfo, cf.File, cf.Line)
//		trace = fmt.Sprintf("%s\n%s", trace, frameInfo)
//	}
//
//	return trace
// }

// todo add trace
func (e *Error) Log() string {
	var code, message, details, trace, sql string
	var err error

	appErr, isAppErr := e.err.(*app.Error)

	if isAppErr {
		err = appErr.UnderlyingErr()
		code = string(appErr.Code)
		message = appErr.Message
		sql, _ = appErr.Data[config.GormSQLKey].(string)

		data, _ := json.Marshal(appErr.Data)
		details = string(data)
	} else {
		data, _ := json.Marshal(e.Response)
		message = string(data)

		dataDetails, _ := json.Marshal(e.Details)
		details = string(dataDetails)
	}

	return fmt.Sprintf(
		"\n\nError: %s\n\nCode: %s\n\nMessage: %s\n\nDetails: %s\n\nStack: \t%s\n\nSql: \t%s\n"+
			"--------------------------------------------------------------------------------------------------------------\n",
		err,
		code,
		message,
		details,
		trace,
		sql,
	)
}

// apiErrorFromAppError transform internal error to an api error
func apiErrorFromAppError(err *app.Error, msg string) *Error {
	apiError := &Error{
		err: err,
	}

	switch err.Code {
	case app.ErrNotFound:
		apiError.Status = 404
		apiError.Message = err.Message
	default:
		apiError.Status = 500
		apiError.Message = "Internal server error"

		dErr, ok := apiErrorFromDomainError(err.Err)
		if ok {
			apiError.Status = dErr.Status
			apiError.Message = dErr.Message
		}
	}

	return apiError
}

func apiErrorFromDomainError(err error) (*Error, bool) {
	apiError := &Error{}
	ok := true

	switch {
	default:
		ok = false
	}

	return apiError, ok
}
