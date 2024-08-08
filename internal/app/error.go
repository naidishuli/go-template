package app

import (
    "errors"
    "fmt"
)

// Error represent an errors happened in the betstion
type Error struct {
    Code    string
    Message string
    Err     error
}

func NewError(code string, errs ...any) error {
    var message string
    var ierr error
    if len(errs) > 0 {
        var ok bool
        err := errs[0]

        ierr, ok = err.(error)
        if ok {
            var appErr *Error
            ok := errors.As(ierr, &appErr)
            if ok {
                return appErr
            }
        }

        message, _ = err.(string)
    }

    return &Error{
        Code:    code,
        Message: message,
        Err:     ierr,
    }
}

func (e *Error) Unwrap() error {
    return e.Err
}

// Error used to implement errors interface
func (e *Error) Error() string {
    return fmt.Sprintf("%s - %s \n %+v", e.Code, e.Message, e.Err)
}
