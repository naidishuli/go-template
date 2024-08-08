package common

import (
    "net/http"

    "github.com/gofiber/fiber/v3"
    "go-template/internal/api/apierror"
    "go-template/internal/app"
    "go-template/pkg/validator"
)

type Validatable interface {
    Validate() error
}

type Processable interface {
    ProcessParams()
}

func ProcessRequest[T any](ctx fiber.Ctx) (*app.Ctx, T, error) {
    var req T
    appCtx := app.NewCtx(ctx.UserContext())

    err := ctx.Bind().URI(&req)
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

func parseMutatorRequest(out any, ctx fiber.Ctx) error {
    err := ctx.Bind().Query(out)
    if err != nil {
        return apierror.BadRequestMalformed(err)
    }

    err = ctx.Bind().Body(out)
    if err != nil {
        return apierror.BadRequestMalformed(err)
    }

    return validate(out)
}

func parseQueryRequest(out any, ctx fiber.Ctx) error {
    err := ctx.Bind().Query(out)
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
    var err error
    v, ok := out.(Validatable)
    if ok {
        err = v.Validate()
    } else {
        err = validator.Get().Struct(out)
    }

    if err != nil {
        return apierror.BadRequest(nil, err)
    }

    return nil
}
