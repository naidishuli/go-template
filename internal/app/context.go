package app

import (
	"context"
	"github.com/hibiken/asynq"
	"go-template/internal/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type Ctx struct {
	ctx context.Context
}

func NewCtx(ctx context.Context) *Ctx {
	return &Ctx{ctx}
}

func (c *Ctx) Context() context.Context {
	return c.ctx
}

// ------------------------------------------------------------------------------------------------

func (c *Ctx) WithAsynqTask(t *asynq.Task) *Ctx {
	return &Ctx{context.WithValue(c.ctx, config.AsynqTaskContextKey, t)}
}

func (c *Ctx) WithDB(db *gorm.DB) *Ctx {
	return &Ctx{context.WithValue(c.ctx, config.GormContextKey, db)}
}

func (c *Ctx) WithDBClauses(clauses ...clause.Expression) *Ctx {
	allClauses := c.DBClauses()
	allClauses = append(allClauses, clauses...)

	return &Ctx{context.WithValue(c.ctx, config.GormClausesContextKey, allClauses)}
}

func (c *Ctx) WithDBScopes(scopes ...func(db *gorm.DB) *gorm.DB) *Ctx {
	allScopes := c.DBScopes()
	allScopes = append(allScopes, scopes...)

	return &Ctx{context.WithValue(c.ctx, config.GormScopesContextKey, allScopes)}
}

func (c *Ctx) WithLogger(log Logger) *Ctx {
	return &Ctx{context.WithValue(c.ctx, config.LoggerContextKey, log)}
}

// ------------------------------------------------------------------------------------------------

func (c *Ctx) DB(defs ...*gorm.DB) *gorm.DB {
	var def *gorm.DB
	if len(defs) > 0 {
		def = defs[0]
	}

	clauses := c.DBClauses()
	scopes := c.DBScopes()

	val := c.ctx.Value(config.GormContextKey)
	if val == nil {
		return def.Clauses(clauses...).Scopes(scopes...)
	}

	db, ok := val.(*gorm.DB)
	if !ok {
		log.Println("context: gorm context key value not as expected (*gorm.DB)")
		return def.Clauses(clauses...).Scopes(scopes...)
	}

	return db.Clauses(clauses...).Scopes(scopes...)
}

func (c *Ctx) DBWithoutScopes(defs ...*gorm.DB) *gorm.DB {
	var def *gorm.DB
	if len(defs) > 0 {
		def = defs[0]
	}

	val := c.ctx.Value(config.GormContextKey)
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

func (c *Ctx) DBClauses() []clause.Expression {
	val := c.ctx.Value(config.GormClausesContextKey)
	if val == nil {
		return []clause.Expression{}
	}

	clauses, ok := val.([]clause.Expression)
	if !ok {
		panic("context: gorm scopes context key value not as expected (func(db *gorm.DB) *gorm.DB)")
	}

	return clauses
}

func (c *Ctx) DBScopes() []func(db *gorm.DB) *gorm.DB {
	val := c.ctx.Value(config.GormScopesContextKey)
	if val == nil {
		return []func(db *gorm.DB) *gorm.DB{}
	}

	scopes, ok := val.([]func(db *gorm.DB) *gorm.DB)
	if !ok {
		panic("context: gorm scopes context key value not as expected (func(db *gorm.DB) *gorm.DB)")
	}

	return scopes
}

func (c *Ctx) Logger(defs ...Logger) Logger {
	var def Logger
	if len(defs) > 0 {
		def = defs[0]
	}

	val := c.ctx.Value(config.LoggerContextKey)
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
