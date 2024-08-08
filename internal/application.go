package internal

import (
    "bets/internal/app"
    "bets/internal/repository"
    "bets/pkg/logger"
    "gorm.io/gorm"
)

type ApplicationConfig struct {
    NoDB    bool
    NoRedis bool
}

type Application struct {
    config ApplicationConfig

    db  *gorm.DB
    log app.Logger

    pkg        *app.Pkg
    repository *app.Repository
    service    *app.Service
    task       *app.Task
    cases      *Cases
}

func Newbetstion(config ApplicationConfig) *Application {
    appl := Application{
        log:    logger.NewLogger(nil, nil),
        config: config,
    }

    appl.db = new(gorm.DB)
    appl.pkg = new(app.Pkg)
    appl.repository = &app.Repository{
        Base: new(repository.Base),
    }
    appl.service = &app.Service{}
    appl.task = &app.Task{}
    appl.cases = &Cases{dep: appl}

    return &appl
}

func (a Application) DB() *gorm.DB {
    return a.db.Session(&gorm.Session{})
}

func (a Application) Log() app.Logger {
    return a.log
}

func (a Application) Pkg() app.Pkg {
    return *a.pkg
}

func (a Application) Repository() app.Repository {
    return *a.repository
}

func (a Application) Service() app.Service {
    return *a.service
}

func (a Application) Task() app.Task {
    return *a.task
}

func (a Application) Cases() app.Cases {
    return *a.cases
}

// Implement CaseInitiator

type Cases struct {
    dep app.App
}
