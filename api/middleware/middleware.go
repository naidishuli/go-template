package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"go-template/internal"
	"go-template/internal/model"
)

type Pkg interface {
	VerifyDataToken(header string, data interface{}) error
}

type Repository interface {
}

type Service interface {
}

type Middleware struct {
	pkg     Pkg
	repo    Repository
	service Service
}

func New(app internal.AppContext) *Middleware {
	return &Middleware{
		pkg:     app.Pkg(),
		repo:    app.Repository(),
		service: app.Service(),
	}
}

func (a *Middleware) Authorize(ctx *fiber.Ctx) error {
	var user model.User
	authHeader := ctx.Get("authorization")

	err := a.pkg.VerifyDataToken(authHeader, &user)
	if err != nil {
		return ctx.Status(401).SendString(err.Error())
	}

	userCtx := context.WithValue(context.Background(), "user", user)
	ctx.SetUserContext(userCtx)
	return ctx.Next()
}
