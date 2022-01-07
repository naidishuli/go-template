package jwt

import "fmt"

type ErrorType string

const (
	MalformedErr    ErrorType = "Malformed token"
	VerificationErr ErrorType = "Error verification method"
	InvalidErr      ErrorType = "Invalid token"
	SignTokenErr    ErrorType = "Error signing token"
)

type tokenError struct {
	reason  ErrorType
	message string
}

func (e tokenError) Error() string {
	return fmt.Sprintf("%s | %s", e.reason, e.message)
}

func malformedError() *tokenError {
	return &tokenError{
		reason:  MalformedErr,
		message: "bearer token malformed or not present",
	}
}

func verificationError() *tokenError {
	return &tokenError{
		reason:  VerificationErr,
		message: "error verification jwtoken method",
	}
}

func invalidError() *tokenError {
	return &tokenError{
		reason:  InvalidErr,
		message: "invalid token",
	}
}

func signError(err error) *tokenError {
	return &tokenError{
		reason:  SignTokenErr,
		message: fmt.Sprintf("cannot sign token string: %s", err),
	}
}
