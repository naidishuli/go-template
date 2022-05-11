package controllers

import (
	"github.com/gofiber/fiber/v2"

	"go-template/internal"
	"go-template/internal/model"
)

type Context struct {
	*internal.Application
}

func NewContext(app *internal.Application) *Context {
	return &Context{app}
}

func (c *Context) App() *internal.Application {
	return c.Application
}

func (c *Context) User(ctx *fiber.Ctx) model.User {
	return ctx.UserContext().Value("user").(model.User)
}
