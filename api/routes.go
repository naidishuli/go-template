package api

import (
	"github.com/gofiber/fiber/v2"
	"go-template/api/common"
	"go-template/api/controller"
	"go-template/api/middleware"
	"go-template/internal/app"
)

// RegisterRoutes used to register api routes to their handlers.
func RegisterRoutes(app app.App, fiberApp *fiber.App) {
	cmnController := common.NewController(app)
	cmnMiddleware := middleware.NewCommon(app)

	privateRoutes := fiberApp.Group("/api", cmnMiddleware.Authorize)

	controller.NewTemp(cmnController).RegisterRoutes(privateRoutes, fiberApp)

	pingRoute(fiberApp)
}

func pingRoute(app *fiber.App) {
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).SendString("ok")
	})
}
