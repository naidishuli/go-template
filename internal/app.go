package internal

import (
	"go-template/internal/app"
	"go-template/internal/config"
	tempRepo "go-template/internal/repository/temp"
	"go-template/internal/services/temp"
	"go-template/pkg/jwt"
	"go-template/pkg/logger"
	"gorm.io/gorm"
)

type Application struct {
	db             *gorm.DB
	pkgPool        *app.Pkg
	repositoryPool *app.Repository
	servicePool    *app.Service
	logger         app.Logger
}

func NewApplication() (*Application, error) {
	appl := Application{
		logger: logger.NewLogger(nil, nil),
	}

	appl.pkgPool = new(app.Pkg)
	appl.repositoryPool = new(app.Repository)
	appl.servicePool = new(app.Service)

	pkgPool, err := initializePkg(&appl)
	if err != nil {
		return nil, err
	}
	*appl.pkgPool = pkgPool

	*appl.repositoryPool = initializeRepository(&appl)
	*appl.servicePool = initializeService(&appl)

	return &appl, nil
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

func initializeRepository(appl *Application) app.Repository {
	return app.Repository{
		Temp: tempRepo.NewTemp(appl),
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
