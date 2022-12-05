package app

import "go-template/pkg/jwt"

//go:generate mockgen -source pkg.go -package app -destination mocks/pkg_mock.go

type Pkg struct {
	JWT
}

type JWT interface {
	VerifyDataToken(header string, data interface{}) error
	GenerateUserToken(userData string) (token jwt.Token, err error)
}
