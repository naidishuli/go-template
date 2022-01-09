package temp

import (
	"github.com/gofiber/fiber/v2"

	"go-template/api/controllers"
	"go-template/api/response"
)

type Controller struct {
	*controllers.Context
}

func NewController(ctx *controllers.Context) Controller {
	return Controller{Context: ctx}
}

// DoSomething does something
// @Tags Temp
// @Summary Just a sample endpoint
// @Description This endpoint just does something.
// @Param id path string true "Just an ID"
// @Param session header string true "JWT"
// @Produce json
// @Success 200 {object} response.StatusOk
// @Failure 401 {object} response.ErrorUnauthorized
// @Failure 404 {object} response.ErrorNotFound
// @Failure 500 {object} response.ErrorInternalServerError
// @Router /temp/{id} [get]
func (c *Controller) DoSomething(ctx *fiber.Ctx) error {
	return ctx.JSON(response.StatusOk{Status: "ok"})
}
