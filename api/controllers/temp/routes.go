package temp

import (
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) RegisterRoutes(privateRoute, publicRoute fiber.Router) {
	tempGroup := privateRoute.Group("/temp")
	tempGroup.Get("/do_something", c.DoSomething)
}
