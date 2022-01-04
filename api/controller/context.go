package controller

import (
	"github.com/gofiber/fiber/v2"

	"go-template/internal"
	"go-template/internal/model"
)

type Context struct {
	app internal.AppContext
}

func NewContext(app internal.AppContext) *Context {
	return &Context{app}
}

func (c *Context) App() internal.AppContext {
	return c.app
}

func (c *Context) User(ctx *fiber.Ctx) model.User {
	return ctx.UserContext().Value("user").(model.User)
}
