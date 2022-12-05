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
func (t *Temp) DoSomethingTemp(arg string, ctx app.Context) error {
	c := app.NewContext(ctx)
	db := c.DB(t.db)
	log := ctx.Log(nil)

	log.Info("do something with the db", db)
	return nil
}
