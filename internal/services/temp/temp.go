package temp

import (
	"go-template/internal/app"
	"gorm.io/gorm"
)

type Temp struct {
	db  *gorm.DB
	log app.Logger

	tempRepo app.TempRepo
}

// NewService create a new services of type Temp.
func NewService(ctx app.App) *Temp {
	return &Temp{
		db:       ctx.DB(),
		log:      ctx.Log(),
		tempRepo: ctx.Repository().Temp,
	}
}

// DoSomething used as an example to follow.
func (t *Temp) DoSomething(ctx *app.Context, arg string) error {
	return t.tempRepo.DoSomethingTemp(ctx, arg)
}

// DoSomethingTransaction used as an example to follow.
func (t *Temp) DoSomethingTransaction(ctx *app.Context, arg string) error {
	iCtx := app.ContextWithDefaults(ctx, app.CtxOption{DB: t.db, Log: t.log})

	return iCtx.DB.Transaction(func(tx *gorm.DB) error {
		// changes to ctx here will not affect the original context passed in parameters
		iCtx.DB = tx
		return t.tempRepo.DoSomethingTemp(iCtx, arg)
	})
}
