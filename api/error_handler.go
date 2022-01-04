package api

import "github.com/gofiber/fiber/v2"

// ErrorHandler is a used as the default error handler
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	return nil
}
