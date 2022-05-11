package repository

import (
	"gorm.io/gorm"
)

type Pool struct {
	*Temp
}

func NewPool(db *gorm.DB) Pool {
	return Pool{
		Temp: NewTemp(db),
	}
}
