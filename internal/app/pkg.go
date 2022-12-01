package app

import "go-template/pkg/jwt"

type Pkg struct {
	JWT
}

type JWT interface {
	VerifyDataToken(header string, data interface{}) error
	GenerateUserToken(userData string) (token jwt.Token, err error)
}
