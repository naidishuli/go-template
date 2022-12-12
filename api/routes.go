package api

import (
	"github.com/gofiber/fiber/v2"
	"go-template/api/common"
	"go-template/api/middlewares"
	"go-template/internal/app"

	"go-template/api/controllers/temp"
)

// RegisterRoutes used to register api routes to their handlers.
func RegisterRoutes(app app.App, fiberApp *fiber.App) {
	ctxController := common.NewContext(app)
	generalMiddlewares := middlewares.NewCommon(app)

	// initialize all controllers here
	tempController := temp.NewController(ctxController)

	// register all the secure routes here
	securedAPI := fiberApp.Group("/api", generalMiddlewares.Authorize)

	tempController.RegisterRoutes(securedAPI, fiberApp)

	pingRoute(fiberApp)
}

func pingRoute(app *fiber.App) {
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("ok")
	})
}
