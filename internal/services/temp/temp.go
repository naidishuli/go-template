package temp

import (
	"go-template/internal/app"
	"gorm.io/gorm"
)

type Temp struct {
	db       *gorm.DB
	tempRepo app.TempRepo
}

// NewService create a new services of type Temp.
func NewService(ctx app.App) *Temp {
	return &Temp{
		db:       ctx.DB(),
		tempRepo: ctx.Repository().Temp,
	}
}

// DoSomething used as an example to follow.
func (t *Temp) DoSomething(arg string, c *app.Context) error {
	return t.tempRepo.DoSomethingTemp(arg, c)
}

// DoSomethingTransaction used as an example to follow.
func (t *Temp) DoSomethingTransaction(arg string, c *app.Context) error {
	ctx := app.NewContext(c)
	db := ctx.DB(t.db)

	return db.Transaction(func(tx *gorm.DB) error {
		ctx.TX = tx
		return t.tempRepo.DoSomethingTemp(arg, ctx)
	})
}
