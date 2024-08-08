package common

import (
    "net/url"

    "bets/internal/app"
    "github.com/gofiber/fiber/v3"
)

type Controller struct {
    app app.App
}

func NewController(app app.App) *Controller {
    return &Controller{app}
}

func (c *Controller) App() app.App {
    return c.app
}

func (c *Controller) RedirectWithData(ctx fiber.Ctx, redirectURL string, data map[string]string) error {
    parsedURL, err := url.Parse(redirectURL)
    if err != nil {
        return err
    }

    params := url.Values{}
    for k, v := range data {
        params.Add(k, v)
    }

    parsedURL.RawQuery = params.Encode()
    return ctx.Redirect().To(parsedURL.String())
}
