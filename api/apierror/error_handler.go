package apierror

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler is a used as the default errors handler
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	apiError, ok := err.(*Error)
	if !ok {
		apiError = New(err)
	}

	return ctx.Status(apiError.Status).JSON(apiError)
}
