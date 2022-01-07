package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type Temp struct {
	db *gorm.DB
}

func NewTemp(db *gorm.DB) *Temp {
	return &Temp{db}
}

// DoSomethingTemp this is an example to follow.
func (t *Temp) DoSomethingTemp(arg string) error {
	fmt.Println(arg)
	return nil
}
