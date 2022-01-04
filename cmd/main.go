package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go-template/api"
	"go-template/internal"
	"go-template/internal/config"
)

func main() {
	application, err := internal.New()
	if err != nil {
		panic(err)
	}

	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: api.ErrorHandler,
	})

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	api.RegisterRoutes(application, fiberApp)
	err = fiberApp.Listen(fmt.Sprintf(":%d", config.Env.Port))
	if err != nil {
		panic(err)
	}
}
