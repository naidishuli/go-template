package repository

import (
	"fmt"
	"go-template/internal/app"
)

type Temp struct {
	repo *repository
}

func NewTemp(repo *repository) *Temp {
	return &Temp{repo}
}

// DoSomethingTemp this is an example to follow.
func (t *Temp) DoSomethingTemp(arg string, ctx *app.Context) error {
	fmt.Println(arg)
	return nil
}
