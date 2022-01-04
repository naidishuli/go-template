package temp

import (
	"go-template/internal/service"
)

type Repository interface {
	DoSomethingTemp()
}

type Temp struct {
	repo Repository
}

// NewTempService create a new service of type Temp
func NewTempService(ctx service.Context) *Temp {
	return &Temp{
		repo: ctx.Repository(),
	}
}

// DoSomething used as an example to follow
func (s *Temp) DoSomething() {
	s.repo.DoSomethingTemp()
}
