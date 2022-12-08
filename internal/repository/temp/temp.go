package temp

import (
	"go-template/internal/app"
	"gorm.io/gorm"
)

type Temp struct {
	db  *gorm.DB
	log app.Logger
}

func NewTemp(app app.App) *Temp {
	return &Temp{
		db:  app.DB(),
		log: app.Log(),
	}
}

// DoSomethingTemp this is an example to follow.
func (t *Temp) DoSomethingTemp(ctx *app.Context, arg string) error {
	iCtx := app.ContextWithDefaults(ctx, app.CtxOption{DB: t.db, Log: t.log})
	iCtx.Log.Info("do something with the db", iCtx.DB)
	return nil
}
