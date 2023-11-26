package service

import "go-template/internal/app"

type Temp struct {
}

func NewTemp(dep app.App) *Temp {
	return &Temp{}
}

func (t Temp) DoSomething(ctx *app.Ctx, arg string) error {
	return nil
}
