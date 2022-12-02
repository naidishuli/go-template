package temp

import (
	"fmt"
	"go-template/internal/app"
	"gorm.io/gorm"
)

type Temp struct {
	db *gorm.DB
}

func NewTemp(db *gorm.DB) *Temp {
	return &Temp{db: db}
}

// DoSomethingTemp this is an example to follow.
func (t *Temp) DoSomethingTemp(arg string, c *app.Context) error {
	ctx := app.NewContext(c)
	db := ctx.DB(t.db)

	fmt.Println("do something with the db", db)
	return nil
}
