package interfaces

import (
	"go-template/internal/pkg"
	"go-template/internal/repository"
)

type ServiceContext interface {
	Pkg() *pkg.Pool
	Repository() *repository.Pool
	ServicePool() interface{}
}
