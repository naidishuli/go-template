package app

import (
	"context"
	"gorm.io/gorm"
)

//go:generate mockgen -source context.go -package mocks -destination mocks/context_mock.go

type Context interface {
	DB(def *gorm.DB) *gorm.DB

	AddValue(key string, value any)
	GetValue(key string) any

	SetTransaction(tx *gorm.DB)
	ClearTransaction()

	SetLogger(log Logger)
	Log(def Logger) Logger
}

// Ctx is passed through every layer of the application
// it will contain dynamic data like:
// db transactions for controlling through layers
// a context logger for individual request/call information
// user information
// etc...
type Ctx struct {
	tx  *gorm.DB // database transaction to be used by the next layer
	ctx context.Context
	log Logger
}

func NewContext(def Context) Context {
	if def != nil {
		return def
	}

	return &Ctx{
		ctx: context.Background(),
	}
}

func (c *Ctx) DB(def *gorm.DB) *gorm.DB {
	if c.tx == nil {
		return def
	}

	return c.tx
}

func (c *Ctx) AddValue(key string, value any) {
	c.ctx = context.WithValue(c.ctx, key, value)
}

func (c *Ctx) GetValue(key string) any {
	return c.ctx.Value(key)
}

func (c *Ctx) SetTransaction(tx *gorm.DB) {
	c.tx = tx
}

func (c *Ctx) ClearTransaction() {
	c.tx = nil
}

func (c *Ctx) SetLogger(log Logger) {
	c.log = log
}

func (c *Ctx) Log(def Logger) Logger {
	if c.log == nil {
		return def
	}

	return c.log
}
