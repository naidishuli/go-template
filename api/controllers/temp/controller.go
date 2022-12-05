package temp

import (
	"github.com/gofiber/fiber/v2"
	"go-template/api/responses"

	"go-template/api/controllers"
)

//go:generate mockgen -source controller.go -package temp -destination mocks/controller_mock.go

type Controller struct {
	*controllers.Common
}

func NewController(ctx *controllers.Common) Controller {
	return Controller{Common: ctx}
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
	return ctx.JSON(responses.StatusOk{Status: "ok"})
}
