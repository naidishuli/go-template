package app

import (
	"context"
	"gorm.io/gorm"
)

//go:generate mockgen -source context.go -package mocks -destination mocks/context_mock.go

// Context is passed through every layer of the application
// it will contain dynamic data like:
// db transactions for controlling through layers
// a context logger for individual request/call information
// user information
// etc...
type Context struct {
	DB  *gorm.DB // database transaction to be used by the next layer
	Log Logger

	ctx context.Context
}

func ContextWithDefaults(c *Context, opts CtxOption) *Context {
	var ctx Context
	if c == nil {
		ctx = Context{
			ctx: context.Background(),
		}
	} else {
		ctx = *c
	}

	if ctx.DB == nil {
		ctx.DB = opts.DB
	}

	if ctx.Log == nil {
		ctx.Log = opts.Log
	}

	return &ctx
}

func (c *Context) AddValue(key string, value any) {
	c.ctx = context.WithValue(c.ctx, key, value)
}

func (c *Context) GetValue(key string) any {
	return c.ctx.Value(key)
}

type CtxOption struct {
	DB  *gorm.DB
	Log Logger
}
