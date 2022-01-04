package internal

import "go-template/internal/pool"

type AppContext interface {
	Pkg() *pool.Pkg
	Repository() *pool.Repository
	Service() *Service
}
