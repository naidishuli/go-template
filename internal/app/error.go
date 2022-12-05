package app

import (
	"fmt"
	"runtime"
)

// ErrCode represents the type of errors in the application level
type ErrCode string

const (
	UndefinedErr ErrCode = "Error.Undefined"

	TokenGenerateErr ErrCode = "Error.Token.Generate"

	DatabaseErr ErrCode = "Error.Repository.Database"
	NotFoundErr ErrCode = "Error.Repository.NotFound"
)

type ErrData map[string]interface{}

// Error represent an errors happened in the application
type Error struct {
	Code    ErrCode
	Message string
	Trace   string

	ErrData

	stackErrData []ErrData
	err          error
}

func NewError(code ErrCode, err error, data ErrData) *Error {
	if err == nil {
		return nil
	}

	appErr, ok := err.(*Error)
	if ok {
		if data != nil {
			appErr.stackErrData = append(appErr.stackErrData, data)
		}
		return appErr
	}

	return &Error{
		Code:    code,
		Message: err.Error(),
		ErrData: data,
		Trace:   trace(),
	}
}

// Error used to implement errors interface
func (e Error) Error() string {
	return e.Message
}

// trace used to return the line of the code that creates the errors
func trace() string {
	pcs := make([]uintptr, 100)
	runtime.Callers(2, pcs)
	callersFrames := runtime.CallersFrames(pcs)

	var trace string
	for {
		cf, more := callersFrames.Next()
		if !more {
			break
		}

		frameInfo := cf.Function
		frameInfo = fmt.Sprintf("%s\n\t%s:%d", frameInfo, cf.File, cf.Line)
		trace = fmt.Sprintf("%s\n%s", trace, frameInfo)
	}

	return trace
}
