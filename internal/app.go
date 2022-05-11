package internal

import (
	"gorm.io/gorm"

	"go-template/internal/pkg"
	"go-template/internal/repository"
	"go-template/internal/services"
)

type Application struct {
	db             *gorm.DB
	pkgPool        *pkg.Pool
	repositoryPool *repository.Pool
	servicePool    *services.Pool
}

func NewApplication() (*Application, error) {
	app := Application{}

	app.pkgPool = new(pkg.Pool)
	app.repositoryPool = new(repository.Pool)
	app.servicePool = new(services.Pool)

	pkgPool, err := pkg.NewPool()
	if err != nil {
		return nil, err
	}
	*app.pkgPool = pkgPool

	repositoryPool := repository.NewPool(app.db)
	*app.repositoryPool = repositoryPool

	servicePool := services.NewPool(&app)
	*app.servicePool = servicePool

	return &app, nil
}

func (a *Application) DB() *gorm.DB {
	return a.db
}

func (a *Application) Pkg() *pkg.Pool {
	return a.pkgPool
}

func (a *Application) Repository() *repository.Pool {
	return a.repositoryPool
}

func (a *Application) Service() *services.Pool {
	return a.servicePool
}

func (a *Application) ServicePool() interface{} {
	return a.servicePool
}
