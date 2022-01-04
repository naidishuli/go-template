package internal

import (
	"gorm.io/gorm"

	"go-template/internal/pool"
	"go-template/internal/service/temp"
)

type Application struct {
	db             *gorm.DB
	pkgPool        *pool.Pkg
	repositoryPool *pool.Repository
	servicePool    *Service
}

func New() (*Application, error) {
	app := Application{}

	app.pkgPool = new(pool.Pkg)
	app.repositoryPool = new(pool.Repository)
	app.servicePool = new(Service)

	pkgPool, err := pool.NewPkg()
	if err != nil {
		return nil, err
	}
	*app.pkgPool = pkgPool

	repositoryPool := pool.NewRepository(app.db)
	*app.repositoryPool = repositoryPool

	servicePool := NewService(&app)
	*app.servicePool = servicePool

	return &app, nil
}

func (a *Application) Pkg() *pool.Pkg {
	return a.pkgPool
}

func (a *Application) Repository() *pool.Repository {
	return a.repositoryPool
}

func (a *Application) Service() *Service {
	return a.servicePool
}

func (a *Application) ServicePool() interface{} {
	return a.servicePool
}

// services

// Service is a pool where all the services will be embedded
type Service struct {
	*temp.Temp
}

func NewService(app *Application) Service {
	return Service{
		Temp: temp.NewTempService(app),
	}
}
