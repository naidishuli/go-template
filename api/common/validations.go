package common

import (
	"github.com/gofiber/fiber/v2"
	"go-template/api/apierror"
	"go-template/internal/app"
	"net/http"
)

type Validatable interface {
	Validate() error
}

type Processable interface {
	ProcessParams()
}

func ParseAndValidate[T any](ctx *fiber.Ctx) (*app.Ctx, T, error) {
	var req T
	appCtx := app.NewCtx(ctx.UserContext())

	err := ctx.ParamsParser(&req)
	if err != nil {
		return appCtx, req, apierror.BadRequestMalformed(err)
	}

	switch ctx.Method() {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		return appCtx, req, parseMutatorRequest(&req, ctx)
	case http.MethodGet:
		return appCtx, req, parseQueryRequest(&req, ctx)
	default:
		return appCtx, req, validate(&req)
	}
}

func parseMutatorRequest(out any, ctx *fiber.Ctx) error {
	err := ctx.BodyParser(out)
	if err != nil {
		return apierror.BadRequestMalformed(err)
	}

	return validate(out)
}

func parseQueryRequest(out any, ctx *fiber.Ctx) error {
	err := ctx.QueryParser(out)
	if err != nil {
		return apierror.BadRequestMalformed(err)
	}

	v, ok := out.(Processable)
	if ok {
		v.ProcessParams()
	}

	return validate(out)
}

func validate(out any) error {
	v, ok := out.(Validatable)
	if ok {
		err := v.Validate()
		if err != nil {
			return apierror.BadRequest(nil, err)
		}
	}

	return nil
}
