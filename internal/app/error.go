package app

import (
	"fmt"
	"runtime"
)

// ErrCode represents the type of errors in the application level
type ErrCode string

const (
	ErrUndefined ErrCode = "Error.Undefined"

	ErrDatabase     ErrCode = "Error.Repository.Database"
	ErrNotFound     ErrCode = "Error.Repository.NotFound"
	ErrDuplicateKey ErrCode = "Error.Repository.DuplicateKey"

	ErrRPCTimeout ErrCode = "Error.RPC.Timeout"
)

type ErrData map[string]interface{}

// Error represent an errors happened in the application
type Error struct {
	Code ErrCode
	Data ErrData

	Err     error
	Message string
}

func NewError(code ErrCode, err error, data ErrData) *Error {
	if err == nil {
		return nil
	}

	appErr, ok := err.(*Error)
	if ok {
		if data != nil {
			for k, v := range data {
				appErr.Data[k] = v
			}
		}
		return appErr
	}

	return &Error{
		Code:    code,
		Message: err.Error(),
		Data:    data,
	}
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) UnderlyingErr() error {
	return e.Err
}

// Error used to implement errors interface
func (e *Error) Error() string {
	return e.Message
}

func NewErrorFromDomain(err error) *Error {
	if err == nil {
		return nil
	}

	switch {
	default:
		return &Error{
			Code:    ErrUndefined,
			Message: "undefined error",
			Err:     err,
		}
	}
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
