package apierror

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
)

// ErrorHandler is a used as the default errors handler
func ErrorHandler(ctx fiber.Ctx, err error) error {
	var fiberErr *fiber.Error
	ok := errors.As(err, &fiberErr)
	if ok {
		return ctx.Status(fiberErr.Code).JSON(fiberErr.Message)
	}

	var apiError *Error
	ok = errors.As(err, &apiError)
	if !ok {
		apiError = New(err)
	}

	// todo a better way to log here
	log.Println(apiError.Log())

	// todo we should put sentry and other log hooks here

	return ctx.Status(apiError.Status).JSON(apiError)
}
