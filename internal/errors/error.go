package errors

import (
	"fmt"
	"runtime"
)

// Code represents the type of errors in the application level
type Code string

const (
	TokenGenerateError Code = "Error.Token.Generate"
	NotFoundError      Code = "Error.NotFound"
)

// Error represent an errors happened in the application
type Error struct {
	Code      Code
	Message   string
	Details   interface{}
	Developer *developer
	Trace     string
}

func New(code Code, err error, details interface{}) error {
	return &Error{
		Code:    code,
		Message: string(code),
		Developer: &developer{
			Details: map[string]interface{}{
				"error":   err,
				"details": details,
			},
		},
		Trace: trace(),
	}
}

// Error used to implement errors interface
func (e Error) Error() string {
	return e.Message
}

// developer contains information of an errors for development purpose only
type developer struct {
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

// trace used to return the line of the code that creates the errors
func trace() string {
	_, fn, line, _ := runtime.Caller(2)
	return fmt.Sprintf("[%s:%d]", fn, line)
}
