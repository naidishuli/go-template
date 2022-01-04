package pool

import (
	"go-template/internal/config"
	"go-template/pkg/jwt"
)

// Pkg used as an initializer of the outside pool module
type Pkg struct {
	*jwt.JWT
}

func NewPkg() (Pkg, error) {
	return Pkg{
		JWT: jwt.New(config.Env.JwtVerificationKey),
	}, nil
}
