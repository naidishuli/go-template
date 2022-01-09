package api

import (
	"github.com/gofiber/fiber/v2"

	"go-template/api/apierror"
	"go-template/internal/config"
)

// ErrorHandler is a used as the default errors handler
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	apiError, ok := err.(*apierror.ApiError)
	if !ok {
		apiError = apierror.New(err)
	}

	if config.Env.IsProduction() {
		apiError.Developer = nil
	}

	return ctx.Status(apiError.Status).JSON(apiError)
}
