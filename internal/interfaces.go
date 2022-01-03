package internal

import (
	"go-template/internal/app/pool"
)

type AppContext interface {
	Pkg() *pool.Pkg
	Repository() *pool.Repository
	Service() *pool.Service
}
