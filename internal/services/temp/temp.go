package temp

import (
	"go-template/internal/services"
)

//go:generate mockgen -source temp.go -package temp -destination temp_mock.go

type Repository interface {
	DoSomethingTemp(string) error
}

type Temp struct {
	repo Repository
}

// NewService create a new services of type Temp.
func NewService(ctx services.Context) *Temp {
	return &Temp{
		repo: ctx.Repository(),
	}
}

// DoSomething used as an example to follow.
func (s *Temp) DoSomething(arg string) error {
	return s.repo.DoSomethingTemp(arg)
}
