package service

import (
	"go-template/internal/pool"
)

type Context interface {
	Pkg() *pool.Pkg
	Repository() *pool.Repository
	ServicePool() interface{}
}
