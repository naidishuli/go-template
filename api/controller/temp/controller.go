package temp

import (
	"github.com/gofiber/fiber/v2"

	"go-template/api/controller"
	"go-template/api/response"
)

type Controller struct {
	*controller.Context
}

func NewController(ctx *controller.Context) Controller {
	return Controller{Context: ctx}
}

// DoSomething does something
// @Tags Temp
// @Summary Just a sample endpoint
// @Description This endpoint just does something.
// @Param id path string true "Just an ID"
// @Param session header string true "JWT"
// @Produce json
// @Success 200 {object} response.SwagStatusOk
// @Router /temp/{id} [get]
func (c *Controller) DoSomething(ctx *fiber.Ctx) error {
	return ctx.JSON(response.SwagStatusOk{Status: "ok"})
}
