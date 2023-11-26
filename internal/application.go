package internal

import (
	"go-template/internal/app"
	"go-template/internal/repository"
	"go-template/internal/service"
	"go-template/pkg/logger"
	"gorm.io/gorm"
)

type ApplicationConfig struct {
	NoDB     bool
	NoRabbit bool
}

type Application struct {
	db             *gorm.DB
	pkgPool        *app.Pkg
	repositoryPool *app.Repository
	servicePool    *app.Service
	logger         app.Logger
}

func NewApplication() *Application {
	appl := Application{
		logger: logger.NewLogger(nil, nil),
	}

	appl.db = new(gorm.DB)
	appl.pkgPool = new(app.Pkg)
	appl.repositoryPool = &app.Repository{
		Temp: new(repository.Temp),
	}
	appl.servicePool = &app.Service{
		Temp: new(service.Temp),
	}

	return &appl
}

func (a Application) DB() *gorm.DB {
	return a.db
}

func (a Application) Pkg() *app.Pkg {
	return a.pkgPool
}

func (a Application) Repository() *app.Repository {
	return a.repositoryPool
}

func (a Application) Service() *app.Service {
	return a.servicePool
}

func (a Application) Log() app.Logger {
	return a.logger
}
