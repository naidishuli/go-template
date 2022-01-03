package pool

import (
	"go-template/internal/service"
	"go-template/internal/service/temp"
)

type TempService temp.Service

type Service struct {
	*TempService
}

func NewService(app service.DepsContext) Service {
	return Service{
		TempService: (*TempService)(temp.NewTempService(app)),
	}
}
