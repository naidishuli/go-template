package services

import (
	"go-template/internal/interfaces"
	"go-template/internal/services/temp"
)

type Pool struct {
	*temp.Temp
}

func NewPool(ctx interfaces.ServiceContext) Pool {
	return Pool{
		Temp: temp.NewService(ctx),
	}
}
