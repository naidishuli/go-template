package api

import (
	"github.com/gofiber/fiber/v2"

	"go-template/api/controllers"
	"go-template/api/controllers/temp"
	"go-template/api/middleware"
	"go-template/internal"
)

// RegisterRoutes used to register api routes to their handlers.
func RegisterRoutes(app *internal.Application, fiberApp *fiber.App) {
	ctxController := controllers.NewContext(app)
	generalMiddlewares := middleware.New(app)

	// initialize all controllers here
	tempController := temp.NewController(ctxController)

	// register all the secure routes here
	securedAPI := fiberApp.Group("/api", generalMiddlewares.Authorize)

	// temp group
	tempGroup := securedAPI.Group("/temp")
	tempGroup.Get("/do_something", tempController.DoSomething)

	// register all the unsecure routes here
	fiberApp.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("ok")
	})
}
