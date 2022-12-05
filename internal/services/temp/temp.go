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
func (t *Temp) DoSomething(arg string, ctx app.Context) error {
	return t.tempRepo.DoSomethingTemp(arg, ctx)
}

// DoSomethingTransaction used as an example to follow.
func (t *Temp) DoSomethingTransaction(arg string, ctx app.Context) error {
	c := app.NewContext(ctx)
	db := c.DB(t.db)

	return db.Transaction(func(tx *gorm.DB) error {
		c.SetTransaction(tx)
		defer c.ClearTransaction()

		return t.tempRepo.DoSomethingTemp(arg, ctx)
	})
}
