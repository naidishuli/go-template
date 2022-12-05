package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-template/api/apierror"
	"time"
)

func New() *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: apierror.ErrorHandler,
	})

	fiberApp.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	fiberApp.Use(logger.New(logger.Config{
		TimeFormat: time.RFC822,
		Format:     "[${time}] - ${ip}:${port} ${status} ${latency} - ${method} ${path}\n",
	}))

	return fiberApp
}
