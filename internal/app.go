package internal

import (
	"go-template/internal/app"
	"go-template/internal/config"
	"go-template/internal/services/temp"
	"go-template/pkg/jwt"
	"gorm.io/gorm"

	"go-template/internal/repository"
)

type Application struct {
	db             *gorm.DB
	pkgPool        *app.Pkg
	repositoryPool *app.Repository
	servicePool    *app.Service
}

func NewApplication() (*Application, error) {
	appl := Application{}

	appl.pkgPool = new(app.Pkg)
	appl.repositoryPool = new(app.Repository)
	appl.servicePool = new(app.Service)

	pkgPool, err := initializePkg(&appl)
	if err != nil {
		return nil, err
	}
	*appl.pkgPool = pkgPool

	*appl.repositoryPool = initializeRepository(appl.db)
	*appl.servicePool = initializeService(&appl)

	return &appl, nil
}

func (a *Application) DB() *gorm.DB {
	return a.db
}

func (a *Application) Pkg() *app.Pkg {
	return a.pkgPool
}

func (a *Application) Repository() *app.Repository {
	return a.repositoryPool
}

func (a *Application) Service() *app.Service {
	return a.servicePool
}

func initializeRepository(db *gorm.DB) app.Repository {
	repo := repository.New(db)

	return app.Repository{
		Temp: repository.NewTemp(repo),
	}
}

func initializeService(appl *Application) app.Service {
	return app.Service{
		Temp: temp.NewService(appl),
	}
}

func initializePkg(appl *Application) (app.Pkg, error) {
	return app.Pkg{
		JWT: jwt.New(config.Env.JwtVerificationKey),
	}, nil
}
