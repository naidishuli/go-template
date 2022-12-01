package temp

import (
	"go-template/internal/app"
)

//go:generate mockgen -source temp.go -package temp -destination temp_mock.go

type Temp struct {
	tempRepo app.TempRepo
}

// NewService create a new services of type Temp.
func NewService(ctx app.App) *Temp {
	return &Temp{
		tempRepo: ctx.Repository().Temp,
	}
}

// DoSomething used as an example to follow.
func (s *Temp) DoSomething(arg string, ctx *app.Context) error {
	return s.tempRepo.DoSomethingTemp(arg, ctx)
}
