package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"go-template/api/apierror"
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

func New(app *internal.Application) *Middleware {
	return &Middleware{
		pkg:     app.Pkg(),
		repo:    app.Repository(),
		service: app.Service(),
	}
}

// Authorize validates the jwt token passed and inject the user data to the request context.
func (a *Middleware) Authorize(ctx *fiber.Ctx) error {
	var user model.User
	authHeader := ctx.Get("authorization")

	err := a.pkg.VerifyDataToken(authHeader, &user)
	if err != nil {
		return apierror.Unauthorized(err, nil)
	}

	//todo add user access, roles etc...

	userCtx := context.WithValue(context.Background(), UserContext, user)
	ctx.SetUserContext(userCtx)
	return ctx.Next()
}
