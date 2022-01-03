package temp

import (
	"go-template/internal/service"
)

type Repository interface {
	DoSomethingTemp()
}

type Service struct {
	repo Repository
}

// NewTempService create a new service of type Temp
func NewTempService(app service.DepsContext) *Service {
	return &Service{
		repo: app.Deps().Repository.(Repository),
	}
}

// DoSomething used as an example to follow
func (s *Service) DoSomething() {
	s.repo.DoSomethingTemp()
}
