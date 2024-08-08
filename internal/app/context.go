package app

import (
    "context"
    "log"

    "go-template/config"
    "gorm.io/gorm"
)

//go:generate mockgen -source context.go -package mocks -destination mocks/context_mock.go

type Context interface {
    context.Context

    WithDB(db *gorm.DB) Context
    WithLogger(log Logger) Context

    DB(defs ...*gorm.DB) *gorm.DB
    Logger(defs ...Logger) Logger
}

type Ctx struct {
    context.Context
}

func NewCtx(ctx context.Context) *Ctx {
    if ctx == nil {
        ctx = context.Background()
    }
    return &Ctx{ctx}
}

func (c *Ctx) WithDB(db *gorm.DB) Context {
    return &Ctx{context.WithValue(c.Context, config.GormContextKey, db)}
}

func (c *Ctx) WithLogger(log Logger) Context {
    return &Ctx{context.WithValue(c.Context, config.LoggerContextKey, log)}
}

// ------------------------------------------------------------------------------------------------

func (c *Ctx) DB(defs ...*gorm.DB) *gorm.DB {
    var def *gorm.DB
    if len(defs) > 0 {
        def = defs[0]
    }

    val := c.Context.Value(config.GormContextKey)
    if val == nil {
        return def
    }

    db, ok := val.(*gorm.DB)
    if !ok {
        log.Println("context: gorm context key value not as expected (*gorm.DB)")
        return def
    }

    return db
}

func (c *Ctx) Logger(defs ...Logger) Logger {
    var def Logger
    if len(defs) > 0 {
        def = defs[0]
    }

    val := c.Context.Value(config.LoggerContextKey)
    if val == nil {
        return def
    }

    logger, ok := val.(Logger)
    if !ok {
        log.Println("context: logger context key value not as expected (app.Logger)")
        return def
    }

    return logger
}
