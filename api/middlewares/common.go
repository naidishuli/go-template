package middlewares

import (
	"context"
	"go-template/api/common"
	"go-template/internal/app"

	"github.com/gofiber/fiber/v2"

	"go-template/api/apierror"
	"go-template/internal/model"
)

//go:generate mockgen -source common.go -package mocks -destination mocks/common_mock.go

type Pkg interface {
	VerifyDataToken(header string, data interface{}) error
}

type Repository interface {
}

type Service interface {
}

type Common struct {
	pkg     Pkg
	repo    Repository
	service Service
}

func NewCommon(app app.App) *Common {
	return &Common{
		pkg:     app.Pkg(),
		repo:    app.Repository(),
		service: app.Service(),
	}
}

// Authorize validates the jwt token passed and inject the user data to the request context.
func (c Common) Authorize(ctx *fiber.Ctx) error {
	var user model.User
	authHeader := ctx.Get("authorization")

	err := c.pkg.VerifyDataToken(authHeader, &user)
	if err != nil {
		return apierror.Unauthorized(err, nil)
	}

	//todo inject to the user its access, roles etc...

	userCtx := context.WithValue(context.Background(), common.UserCtx, user)
	ctx.SetUserContext(userCtx)
	return ctx.Next()
}
