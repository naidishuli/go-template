package orm

import (
	"go-template/internal/app"
	"gorm.io/gorm"
)

type Base[T any] struct {
	db *gorm.DB
}

func (b Base[T]) Create(ctx *app.Ctx, value any) error {
	db := ctx.DB(b.db)
	rtx := db.Create(value)
	return app.GormErr(rtx, nil)
}

func (b Base) Transaction(ctx *app.Ctx, fun func(ctx *app.Ctx) error) error {
	db := ctx.DB(b.db)
	return db.Transaction(func(tx *gorm.DB) error {
		tCtx := ctx.WithDB(tx)
		return fun(tCtx)
	})
}
