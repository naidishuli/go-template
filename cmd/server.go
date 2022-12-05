package main

import (
	"fmt"
	"go-template/api/apierror"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go-template/api"
	"go-template/internal"
	"go-template/internal/config"
)

func main() {
	app, err := internal.NewApplication()
	if err != nil {
		panic(err)
	}

	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: apierror.ErrorHandler,
	})

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	api.RegisterRoutes(app, fiberApp)
	err = fiberApp.Listen(fmt.Sprintf(":%d", config.Env.Port))
	if err != nil {
		panic(err)
	}
}
