package jwt

import (
	"errors"
)

type ErrorReason string

var (
	ErrTokenMalformed = errors.New("token is malformed")
	ErrTokenInvalid   = errors.New("token is invalid")
	ErrTokenSign      = errors.New("token sign is not correct")
	ErrVerification   = errors.New("verification error")
)
