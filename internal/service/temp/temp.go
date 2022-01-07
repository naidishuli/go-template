package temp

import (
	"go-template/internal/service"
)

//go:generate mockgen -source temp.go -package temp -destination temp_mock.go

type Repository interface {
	DoSomethingTemp(string) error
}

type Temp struct {
	repo Repository
}

// NewTempService create a new service of type Temp.
func NewTempService(ctx service.Context) *Temp {
	return &Temp{
		repo: ctx.Repository(),
	}
}

// DoSomething used as an example to follow.
func (s *Temp) DoSomething(arg string) error {
	return s.repo.DoSomethingTemp(arg)
}
