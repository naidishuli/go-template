package app

import (
	"gorm.io/gorm"

	"go-template/internal/app/pool"
	"go-template/internal/service"
)

type Application struct {
	db             *gorm.DB
	pkgPool        *pool.Pkg
	repositoryPool *pool.Repository
	servicePool    *pool.Service
}

func New() (*Application, error) {
	app := Application{}

	app.pkgPool = new(pool.Pkg)
	app.repositoryPool = new(pool.Repository)
	app.servicePool = new(pool.Service)

	pkgPool, err := pool.NewPkg()
	if err != nil {
		return nil, err
	}
	*app.pkgPool = pkgPool

	repositoryPool := pool.NewRepository(app.db)
	*app.repositoryPool = repositoryPool

	servicePool := pool.NewService(&app)
	*app.servicePool = servicePool

	return &app, nil
}

func (a *Application) Pkg() *pool.Pkg {
	return a.pkgPool
}

func (a *Application) Repository() *pool.Repository {
	return a.repositoryPool
}

func (a *Application) Service() *pool.Service {
	return a.servicePool
}

func (a *Application) Deps() service.Deps {
	return service.Deps{
		Pkg:        a.pkgPool,
		Repository: a.repositoryPool,
		Service:    a.servicePool,
	}
}
