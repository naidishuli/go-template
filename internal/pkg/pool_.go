package pkg

import (
	"go-template/internal/config"
	"go-template/pkg/jwt"
)

type Pool struct {
	*jwt.JWT
}

func NewPool() (Pool, error) {
	return Pool{
		JWT: jwt.New(config.Env.JwtVerificationKey),
	}, nil
}
