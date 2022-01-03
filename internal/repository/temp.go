package repository

import "gorm.io/gorm"

type Temp struct {
	db *gorm.DB
}

func NewTemp(db *gorm.DB) *Temp {
	return &Temp{db}
}

// DoSomethingTemp this is an example to follow
func (t *Temp) DoSomethingTemp() {

}
