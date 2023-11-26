package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-template/api/common"
)

type Temp struct {
	*common.Controller
}

func NewTemp(ctr *common.Controller) *Temp {
	return &Temp{
		Controller: ctr,
	}
}

func (t Temp) RegisterRoutes(private, public fiber.Router) {

}
