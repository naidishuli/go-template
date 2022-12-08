package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-template/api"
	"go-template/internal/app"

	"go-template/internal/model"
)

type Common struct {
	app app.App
}

func NewContext(app app.App) *Common {
	return &Common{app}
}

func (c Common) App() app.App {
	return c.app
}

func (c Common) User(ctx *fiber.Ctx) model.User {
	return ctx.UserContext().Value(api.UserCtx).(model.User)
}

func (c Common) UserAccess(ctx *fiber.Ctx) (model.User, app.UserAccess, error) {
	user := c.User(ctx)

	//todo here we can get the access injected to the user

	//access, ok := user.Access.(interfaces.Access)
	//if !ok {
	//	err := errors.New("access interface conversion error")
	//	c.log.Errorf("%+v\n", err)
	//	return user, nil, apierror.InternalServerError(err)
	//}

	return user, nil, nil
}
